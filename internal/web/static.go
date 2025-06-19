package web

import (
	"embed"
	"io/fs"
	"net/http"
	"path"
	"strings"
)

//go:embed dist/*
var embeddedFiles embed.FS

func RegisterStaticRoutes(mux *http.ServeMux) {
	distFS, err := fs.Sub(embeddedFiles, "dist")
	if err != nil {
		panic("failed to setup embedded filesystem: " + err.Error())
	}
	fileServer := http.FileServer(http.FS(distFS))

	mux.Handle("/assets/", fileServer)
	mux.Handle("/favicon.ico", fileServer)

	// Catch-all for SPA routing — always serve index.html
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fPath := strings.TrimPrefix(path.Clean(r.URL.Path), "/")
		if fPath == "" {
			fPath = "index.html"
		}

		f, err := distFS.Open(fPath)
		if err == nil {
			defer f.Close()
			// Use FileServer if it's a valid file
			fileServer.ServeHTTP(w, r)
			return
		}

		data, err := fs.ReadFile(distFS, "index.html")
		if err != nil {
			http.Error(w, "index.html not found", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})
}
