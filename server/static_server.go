package main

import (
	"flag"
	"github.com/boomka/cv/internal/handlers"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	baseDir := flag.String("base", "", "Base directory to look for files, default uses current directory")
	fullBaseDir, _ := filepath.Abs(*baseDir)

	mux := http.NewServeMux()
	frontend := handlers.Gzip(handlers.NewFrontendHandler(fullBaseDir))

	mux.Handle("/", frontend)

	l := "127.0.0.1:8877"
	log.Printf("Starting HTTP Server at %q", l)
	log.Fatal(http.ListenAndServe(l, mux))
}
