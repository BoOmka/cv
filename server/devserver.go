//go:build ignore
// +build ignore

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/vugu/vugu/simplehttp"
)

func main() {
	wd, _ := os.Getwd()
	l := "127.0.0.1:19944"
	log.Printf("Starting HTTP Server at %q", l)
	h := simplehttp.New(wd+"/app", false)
	log.Fatal(http.ListenAndServe(l, h))
}
