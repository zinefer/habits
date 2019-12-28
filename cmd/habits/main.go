package main

import (
	"log"
	"os"

	"github.com/zinefer/habits/cmd/habits/serve"
)

var subcommand string

func main() {
	if len(os.Args) > 1 {
		subcommand = os.Args[1]
	}

	// Remove subcommand attr
	os.Args = append(os.Args[:1], os.Args[2:]...)

	switch subcommand {
	case "serve":
		serve.Run()
	default:
		log.Fatal("Available commands: serve")
	}
}
