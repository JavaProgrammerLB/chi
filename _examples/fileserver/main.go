//
// FileServer
// ===========
// This example demonstrates how to serve static files from your filesystem.
//
//
// Boot the server:
// ----------------
// $ go run main.go
//
// Client requests:
// ----------------
// $ curl http://localhost:3333/files/
// <pre>
// <a href="notes.txt">notes.txt</a>
// </pre>
//
// $ curl http://localhost:3333/files/notes.txt
// Notessszzz
//
package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// (1) chi.NewRouter()
	r := chi.NewRouter()

	// Mux里的Use函数
	r.Use(middleware.Logger)

	// Index handler
	// Get函数
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		//ResponseWriter的Write函数，[]byte("hi")
		w.Write([]byte("hi"))
	})

	// Create a route along /files that will serve contents from
	// the ./data/ folder.
	// os包里的Getwd()函数
	workDir, _ := os.Getwd()
	// http包里的Dir()函数
	filesDir := http.Dir(filepath.Join(workDir, "data"))
	// 调用本地的FileServer函数
	FileServer(r, "/files", filesDir)

	http.ListenAndServe(":3333", r)
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
