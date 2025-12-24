package main

import (
	"embed"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	ttpb "github.com/chrisfenner/tpm.tools/proto"
	"google.golang.org/protobuf/proto"

	"github.com/chrisfenner/tpm.tools/pkg/httphelpers"
)

//go:embed generated/*.html
var generatedHTML embed.FS

//go:embed dist/bundle.js
var generatedJS embed.FS

//go:embed statics/*
var statics embed.FS

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}

	http.HandleFunc("/manifest.json", httphelpers.StaticallyServe(statics, "statics/manifest.json", "application/manifest+json"))
	http.HandleFunc("/media/favicon-192.png", httphelpers.StaticallyServe(statics, "statics/media/favicon-192.png", "image/png"))
	http.HandleFunc("/media/favicon-512.png", httphelpers.StaticallyServe(statics, "statics/media/favicon-512.png", "image/png"))
	http.HandleFunc("/media/github-mark-white.svg", httphelpers.StaticallyServe(statics, "statics/media/github-mark-white.svg", "image/svg+xml"))
	http.HandleFunc("/", httphelpers.StaticallyServe(generatedHTML, "generated/index.html", "text/html"))
	http.HandleFunc("/rc", httphelpers.StaticallyServe(generatedHTML, "generated/rc.html", "text/html"))

	// js assets
	http.HandleFunc("/bundle.js", httphelpers.StaticallyServe(generatedJS, "dist/bundle.js", "text/javascript"))

	// REST APIs for various site features
	http.HandleFunc("/rc/lookup", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "only POST supported", http.StatusMethodNotAllowed)
			return
		}
		var req ttpb.ReturnCodeLookupRequest
		reqData, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read request", http.StatusInternalServerError)
			return
		}
		if err := proto.Unmarshal(reqData, &req); err != nil {
			http.Error(w, "failed to unmarshal request", http.StatusInternalServerError)
			return
		}

		// TODO: Actual business logic
		// For now, try to parse the query for a hex number.
		// Any error, just use 0.
		var queryValue int64
		queryValue, err = strconv.ParseInt(req.Query, 16, 32)
		if err != nil {
			queryValue = 0
		}

		rsp := ttpb.ReturnCodeLookupResponse{
			Result: []*ttpb.ReturnCodeLookupResult{
				&ttpb.ReturnCodeLookupResult{
					Name:  "TEST",
					Value: int32(queryValue),
				},
			},
		}

		rspData, err := proto.Marshal(&rsp)
		if err != nil {
			http.Error(w, "failed to marshal response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/x-protobuf")
		w.Write(rspData)
	})

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
