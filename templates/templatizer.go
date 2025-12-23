// Small Go program to render all the static HTML files based on templatized HTML.
package main

import (
	"flag"
	"html/template"
	"log"
	"os"
	"path"
	"strings"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
)

var (
	templateDir = flag.String("template_dir", "templates", "input directory for the HTML templates")
	outputDir   = flag.String("output_dir", "generated", "output directory for the generated files")
)

func main() {
	flag.Parse()

	temps, err := os.ReadDir(*templateDir)

	if err != nil {
		log.Fatalf("Could not read templates from %q: %v", *templateDir, err)
	}

	var successFiles []string

	allTemplates, err := template.ParseGlob(path.Join(*templateDir, "internal", "*.html.tmpl"))
	if err != nil {
		log.Fatalf("Could not parse internal templates: %v", err)
	}
	allTemplates, err = allTemplates.ParseGlob(path.Join(*templateDir, "*.html.tmpl"))
	if err != nil {
		log.Fatalf("Could not parse templates: %v", err)
	}

	// Initialize the HTML minifier.
	m := minify.New()
	m.AddFunc("text/html", html.Minify)

	// At startup, render each output template.
	for _, temp := range temps {
		if !strings.HasSuffix(temp.Name(), ".html.tmpl") {
			continue
		}
		genFileName := path.Join(*outputDir, strings.ReplaceAll(temp.Name(), ".html.tmpl", ".html"))
		w, err := os.Create(genFileName)
		if err != nil {
			log.Fatalf("Could not create file %q for writing: %v", genFileName, err)
		}
		defer w.Close()

		// Pass writes through a minifying writer for auto-formatting.
		mw := m.Writer("text/html", w)
		defer mw.Close()

		if err := allTemplates.ExecuteTemplate(mw, temp.Name(), nil); err != nil {
			log.Fatalf("Could not execute template for %q: %v", genFileName, err)
		}

		successFiles = append(successFiles, genFileName)
	}
	log.Printf("Generated the following %d files:\n  %s", len(successFiles), strings.Join(successFiles, "\n  "))
}
