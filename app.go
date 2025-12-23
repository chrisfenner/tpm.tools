package main

import (
	"embed"
	"log"
	"net/http"
	"os"

	"github.com/chrisfenner/tpm.tools/pkg/httphelpers"
)

//go:embed generated/*.html
var generatedHTML embed.FS

//go:embed statics/*
var statics embed.FS

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}

	http.HandleFunc("/styles.css", httphelpers.StaticallyServe(statics, "statics/styles.css", "text/css"))
	http.HandleFunc("/media/logo-small.png", httphelpers.StaticallyServe(statics, "statics/media/logo-small.png", "image/png"))
	http.HandleFunc("/manifest.json", httphelpers.StaticallyServe(statics, "statics/manifest.json", "application/manifest+json"))
	http.HandleFunc("/media/favicon-192.png", httphelpers.StaticallyServe(statics, "statics/media/favicon-192.png", "image/png"))
	http.HandleFunc("/media/favicon-512.png", httphelpers.StaticallyServe(statics, "statics/media/favicon-512.png", "image/png"))
	http.HandleFunc("/media/github-mark-white.svg", httphelpers.StaticallyServe(statics, "statics/media/github-mark-white.svg", "image/svg+xml"))
	http.HandleFunc("/", httphelpers.StaticallyServe(generatedHTML, "generated/index.html", "text/html"))
	http.HandleFunc("/rc", httphelpers.StaticallyServe(generatedHTML, "generated/rc.html", "text/html"))

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
