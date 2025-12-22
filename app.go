package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.Redirect(w, r, "/index.html", http.StatusMovedPermanently)
	// })

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
