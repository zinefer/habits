package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

var listenAddr string

func main() {
	flag.StringVar(&listenAddr, "listen-addr", ":3000", "server listen address")
	flag.Parse()

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	fmt.Printf("Listening on %s\n", listenAddr)
	http.ListenAndServe(listenAddr, r)
}
