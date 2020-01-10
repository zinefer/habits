package serve

import (
	"crypto/tls"
	"encoding/gob"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/acme/autocert"

	"github.com/jmoiron/sqlx"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/sessions"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"

	"github.com/zinefer/habits/internal/pkg/database/manager"
	"github.com/zinefer/habits/internal/pkg/database/migrator"
	"github.com/zinefer/habits/internal/pkg/subcommander"

	"github.com/zinefer/habits/internal/habits/config"
	"github.com/zinefer/habits/internal/habits/config/routes"
	"github.com/zinefer/habits/internal/habits/middlewares/database"
	sessionMW "github.com/zinefer/habits/internal/habits/middlewares/session"
	"github.com/zinefer/habits/internal/habits/models/user"
)

var (
	configuration *config.Configuration
	session       *sessions.CookieStore
	db            *sqlx.DB
)

func init() {
	gob.Register(&user.User{})
}

// Subcommand for the serve task
type Subcommand struct {
	config *config.Configuration
	db     *sqlx.DB
}

// New serve subcommand
func New(config *config.Configuration, db *sqlx.DB) *Subcommand {
	return &Subcommand{
		config: config,
		db:     db,
	}
}

// Subcommander configures the subcommander instance for this subtask
func (*Subcommand) Subcommander() *subcommander.Subcommander {
	return subcommander.New()
}

// Run the http server
func (c *Subcommand) Run() bool {
	session = sessions.NewCookieStore([]byte(c.config.SessionSecret))
	gothic.Store = session

	if c.config.IsProduction() {
		migrate := migrator.New(c.db, config.DatabaseMigrationPath)
		if migrate.MigrationsTableExists() {
			fmt.Println("Executing migrations if they exist")
			err := migrate.Migrate()
			if err != nil {
				fmt.Println("Fatal Error: unable to migrate database")
				panic(err)
			}
		} else {
			fmt.Println("Fresh database detected, loading schema")
			dbMan := manager.New(c.db)
			err := dbMan.Load(config.DatabaseSchemaPath)
			if err != nil {
				fmt.Println("Fatal Error: unable to load database schema")
				panic(err)
			}
		}
	}

	baseAuthURL := "https://" + c.config.Hostname + "/api/auth/"
	oauthProviders := []goth.Provider{}
	if len(c.config.GithubClientID) > 0 && len(c.config.GithubClientSecret) > 0 {
		oauthProviders = append(oauthProviders, github.New(c.config.GithubClientID, c.config.GithubClientSecret, baseAuthURL+"github/callback"))
	}

	if len(c.config.FacebookClientID) > 0 && len(c.config.FacebookClientSecret) > 0 {
		oauthProviders = append(oauthProviders, facebook.New(c.config.FacebookClientID, c.config.FacebookClientSecret, baseAuthURL+"facebook/callback"))
	}

	if len(c.config.GoogleClientID) > 0 && len(c.config.GoogleClientSecret) > 0 {
		oauthProviders = append(oauthProviders, google.New(c.config.GoogleClientID, c.config.GoogleClientSecret, baseAuthURL+"google/callback"))
	}
	goth.UseProviders(oauthProviders...)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(database.DbContextMiddleware(c.db))
	r.Use(sessionMW.SessionContextMiddleware(session))
	r.Use(rerouteToIndexOn404)

	routes.Define(r)

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "web/dist")
	FileServer(r, "/", filesDir)

	var certMan *autocert.Manager
	if c.config.IsProduction() {
		certMan = &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			Cache:      autocert.DirCache(config.CertsConfigPath),
			HostPolicy: autocert.HostWhitelist(c.config.Hostname, "www."+c.config.Hostname),
		}

		tlsServer := &http.Server{
			Addr:    ":https",
			Handler: r,
			TLSConfig: &tls.Config{
				GetCertificate: certMan.GetCertificate,
			},
		}

		go func() {
			fmt.Printf("Starting HTTPS server on %s\n", tlsServer.Addr)
			err := tlsServer.ListenAndServeTLS("", "")
			if err != nil {
				fmt.Println("tlsServer.ListendAndServeTLS() failed with")
				panic(err)
			}
		}()
	}

	server := &http.Server{
		Addr:    c.config.ListenAddress,
		Handler: r,
	}

	if c.config.IsProduction() {
		server.Handler = certMan.HTTPHandler(nil)
	}

	fmt.Printf("Starting HTTP server on %s\n", c.config.ListenAddress)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("server.ListenAndServe() failed with")
		panic(err)
	}
	return true
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, filesDir string) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	root := http.Dir(filesDir)
	fs := http.StripPrefix(path, http.FileServer(root))
	err := filepath.Walk(filesDir, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}

		path, err = filepath.Rel(filesDir, path)
		if err != nil {
			panic(err)
		}

		r.Get("/"+path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fs.ServeHTTP(w, r)
		}))

		return err
	})

	if err != nil {
		panic(err)
	}

	r.Get("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}

func rerouteToIndexOn404(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		routePath := rctx.RoutePath
		if routePath == "" {
			if r.URL.RawPath != "" {
				routePath = r.URL.RawPath
			} else {
				routePath = r.URL.Path
			}
		}

		tctx := chi.NewRouteContext()
		if !rctx.Routes.Match(tctx, r.Method, routePath) {
			r.URL.Path = "/"
			rctx.RoutePath = "/"
		}

		next.ServeHTTP(w, r)
	})
}
