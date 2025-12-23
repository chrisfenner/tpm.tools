package httphelpers

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
)

// Statically serves the contents of the given file with the given content-type.
// Panics if there is an error reading the file from the file system.
func StaticallyServe(fs fs.FS, srcPath string, contentType string) func(w http.ResponseWriter, r *http.Request) {
	f, err := fs.Open(srcPath)
	if err != nil {
		panic(fmt.Sprintf("could not open %q: %v", srcPath, err))
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		panic(fmt.Sprintf("could not read %q: %v", srcPath, err))
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		w.Write(data)
	}
}
