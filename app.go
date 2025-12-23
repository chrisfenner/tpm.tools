package main

import (
	"embed"
	"log"
	"net/http"
	"os"
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

	http.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) {
		data, err := statics.ReadFile("statics/styles.css")
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "text/css")
		w.Write(data)
	})

	http.HandleFunc("/media/logo-small.png", func(w http.ResponseWriter, r *http.Request) {
		data, err := statics.ReadFile("statics/media/logo-small.png")
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "image/png")
		w.Write(data)
	})

	http.HandleFunc("/manifest.json", func(w http.ResponseWriter, r *http.Request) {
		data, err := statics.ReadFile("statics/manifest.json")
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/manifest+json")
		w.Write(data)
	})

	http.HandleFunc("/media/favicon-192.png", func(w http.ResponseWriter, r *http.Request) {
		data, err := statics.ReadFile("statics/media/favicon-192.png")
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "image/png")
		w.Write(data)
	})

	http.HandleFunc("/media/favicon-512.png", func(w http.ResponseWriter, r *http.Request) {
		data, err := statics.ReadFile("statics/media/favicon-512.png")
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "image/png")
		w.Write(data)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := generatedHTML.ReadFile("generated/index.html")
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write(data)
	})

	http.HandleFunc("/rc", func(w http.ResponseWriter, r *http.Request) {
		data, err := generatedHTML.ReadFile("generated/rc.html")
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write(data)
	})

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
