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
	// Postgres driver
	_ "github.com/lib/pq"

	"github.com/zinefer/habits/internal/habits/config"
	"github.com/zinefer/habits/internal/habits/controllers/auth"
	"github.com/zinefer/habits/internal/habits/middlewares/authorize"
	"github.com/zinefer/habits/internal/habits/middlewares/database"
	sessionMW "github.com/zinefer/habits/internal/habits/middlewares/session"
	"github.com/zinefer/habits/internal/habits/models/user"
)

var (
	configuration *config.Configuration
	session       *sessions.CookieStore
	db            *sqlx.DB
)

var schema = `
CREATE TABLE users (
    name text,
    nickname text,
	email text,
	provider text
);`

func init() {
	gob.Register(&user.User{})
}

// Run the http server
func Run() {
	configuration = config.New()

	session = sessions.NewCookieStore([]byte(configuration.SessionSecret))

	gothic.Store = session

	goth.UseProviders(
		github.New(configuration.GithubClientID, configuration.GithubClientSecret, "http://127.0.0.1:3000/auth/github/callback"),
		//google.New(configuration.GoogleClientID, configuration.GoogleClientSecret, "http://localhost:3000/auth/google/callback"),
		//facebook.New(configuration.FacebookClientID, configuration.FacebookClientSecret, "http://localhost:3000/auth/facebook/callback"),
	)

	db, err := sqlx.Open("postgres", "host=127.0.0.1 user=postgres")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//db.MustExec(schema)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(database.DbContextMiddleware(db))
	r.Use(sessionMW.SessionContextMiddleware(session))

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "web/dist")
	FileServer(r, "/", http.Dir(filesDir))

	r.Get("/auth/{provider}/callback", auth.Callback())
	r.Get("/auth/{provider}", auth.SignIn())
	r.Get("/logout", auth.SignOut())

	r.Route("/api", func(r chi.Router) {
		r.Use(authorize.AuthorizeMiddleware())

		r.Get("/test", func(res http.ResponseWriter, req *http.Request) {
			res.Write([]byte(fmt.Sprintf("hello")))
		})
	})

	fmt.Printf("Listening on %s\n", configuration.ListenAddress)
	http.ListenAndServe(configuration.ListenAddress, r)
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
