package main

import (
	"embed"
	"io"
	"log"
	"maps"
	"net/http"
	"os"
	"slices"
	"strconv"

	ttpb "github.com/chrisfenner/tpm.tools/proto"
	"google.golang.org/protobuf/proto"

	"github.com/chrisfenner/tpm.tools/pkg/httphelpers"
	"github.com/chrisfenner/tpm.tools/pkg/jsonproto"
	"github.com/chrisfenner/tpm.tools/pkg/rc"
)

//go:embed generated/*.html
var generatedHTML embed.FS

//go:embed dist/bundle.js
var generatedJS embed.FS

//go:embed statics/*
var statics embed.FS

//go:embed data/part3.json
var part3json []byte

func main() {
	// Load the json/proto data
	cmdLookupMap, err := jsonproto.LoadCommandProtos(part3json)
	if err != nil {
		log.Fatalf("could not load command json: %v", err)
	}
	cmdNames := slices.Collect(maps.Keys(cmdLookupMap))
	slices.Sort(cmdNames)

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
	http.HandleFunc("/cmd", httphelpers.StaticallyServe(generatedHTML, "generated/cmd.html", "text/html"))

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

		rsp := ttpb.ReturnCodeLookupResponse{}

		var queryValue int64
		queryValue, err = strconv.ParseInt(req.Query, 16, 32)
		// If we were able to parse the input as a hex integer, move right along.
		if err == nil {
			results, err := rc.LookupResponseCodeByValue(int32(queryValue))
			// We don't actually care about errors looking up the RC.
			if err == nil && len(results) > 0 {
				rsp.Result = results
			}
		}

		rspData, err := proto.Marshal(&rsp)
		if err != nil {
			http.Error(w, "failed to marshal response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/x-protobuf")
		w.Write(rspData)
	})

	http.HandleFunc("/cmd/list", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "only POST supported", http.StatusMethodNotAllowed)
			return
		}
		var req ttpb.GetAllCommandNamesRequest
		reqData, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read request", http.StatusInternalServerError)
			return
		}
		if err := proto.Unmarshal(reqData, &req); err != nil {
			http.Error(w, "failed to unmarshal request", http.StatusInternalServerError)
			return
		}

		rsp := ttpb.GetAllCommandNamesResponse{
			Name: cmdNames,
		}

		rspData, err := proto.Marshal(&rsp)
		if err != nil {
			http.Error(w, "failed to marshal response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/x-protobuf")
		w.Write(rspData)
	})

	http.HandleFunc("/cmd/lookup", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "only POST supported", http.StatusMethodNotAllowed)
			return
		}
		var req ttpb.GetCommandDescriptionRequest
		reqData, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read request", http.StatusInternalServerError)
			return
		}
		if err := proto.Unmarshal(reqData, &req); err != nil {
			http.Error(w, "failed to unmarshal request", http.StatusInternalServerError)
			return
		}

		rsp := ttpb.GetCommandDescriptionResponse{}

		data, ok := cmdLookupMap[req.Name]
		if ok {
			rsp.Description = data
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
