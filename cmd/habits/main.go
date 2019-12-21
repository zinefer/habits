package main

import (
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
	_ "github.com/lib/pq"

	"github.com/zinefer/habits/internal/habits/config"
	"github.com/zinefer/habits/internal/habits/controllers/auth"
	"github.com/zinefer/habits/internal/habits/middlewares"
)

var (
	configuration *config.Configuration
	db            *sqlx.DB
)

var schema = `
CREATE TABLE users (
    name text,
    nickname text,
	email text,
	provider text
);`

func main() {
	configuration = config.New()

	gothic.Store = sessions.NewCookieStore([]byte(configuration.SessionSecret))

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

	db.MustExec(schema)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middlewares.DbContextMiddleware(db))

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "web/dist")
	FileServer(r, "/", http.Dir(filesDir))

	r.Get("/auth/{provider}/callback", auth.Callback())
	r.Get("/auth/{provider}", auth.SignIn())
	r.Get("/logout/{provider}", auth.SignOut())

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
