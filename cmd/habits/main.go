package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/zinefer/habits/internal/habits/config"
)

var (
	configuration *config.Configuration
)

func main() {
	configuration = config.New()
	flag.Parse()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
		w.Write([]byte("welcome"))
	})*/

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "web/dist")
	FileServer(r, "/", http.Dir(filesDir))

	fmt.Printf("Listening on %s\n", listenAddr)
	http.ListenAndServe(listenAddr, r)
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
