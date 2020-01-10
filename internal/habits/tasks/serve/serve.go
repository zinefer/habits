package serve

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/sessions"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"

	//"github.com/markbates/goth/providers/google"
	//"github.com/markbates/goth/providers/facebook"

	"github.com/jmoiron/sqlx"

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

	goth.UseProviders(
		github.New(c.config.GithubClientID, c.config.GithubClientSecret, "http://127.0.0.1:3000/api/auth/github/callback"),
		//google.New(configuration.GoogleClientID, configuration.GoogleClientSecret, "http://localhost:3000/api/auth/google/callback"),
		//facebook.New(configuration.FacebookClientID, configuration.FacebookClientSecret, "http://localhost:3000/api/auth/facebook/callback"),
	)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(database.DbContextMiddleware(c.db))
	r.Use(sessionMW.SessionContextMiddleware(session))
	r.Use(rerouteToIndexOn404)

	routes.Define(r)

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "web/dist")
	FileServer(r, "/", filesDir)

	fmt.Printf("Listening on %s\n", c.config.ListenAddress)
	http.ListenAndServe(c.config.ListenAddress, r)
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
