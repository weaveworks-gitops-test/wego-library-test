package main

import (
	"context"
	"embed"
	"io/fs"
	"net/http"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/weaveworks/weave-gitops/pkg/server"
)

var addr = "0.0.0.0:8000"
var log = logrus.New()

func main() {
	ctx := context.Background()

	mux := http.NewServeMux()

	cfg, err := server.DefaultConfig()
	if err != nil {
		panic(err)
	}

	handler, err := server.NewApplicationsHandler(ctx, cfg)
	if err != nil {
		panic(err)
	}

	assetFS := getAssets()
	assetHandler := http.FileServer(http.FS(assetFS))
	redirector := createRedirector(assetFS, log)

	mux.Handle("/v1/", handler)
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Assume anything with a file extension in the name is a static asset.
		extension := filepath.Ext(req.URL.Path)
		// We use the golang http.FileServer for static file requests.
		// This will return a 404 on normal page requests, ie /some-page.
		// Redirect all non-file requests to index.html, where the JS routing will take over.
		if extension == "" {
			redirector(w, req)
			return
		}
		assetHandler.ServeHTTP(w, req)
	}))

	log.Info("starting server on " + addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		panic(err)
	}

}

//go:embed dist/*
var static embed.FS

func getAssets() fs.FS {
	f, err := fs.Sub(static, "dist")

	if err != nil {
		panic(err)
	}
	return f
}

func createRedirector(fsys fs.FS, log logrus.FieldLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		indexPage, err := fsys.Open("index.html")

		if err != nil {
			log.Error(err, "could not open index.html page")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		stat, err := indexPage.Stat()
		if err != nil {
			log.Error(err, "could not get index.html stat")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		bt := make([]byte, stat.Size())
		_, err = indexPage.Read(bt)

		if err != nil {
			log.Error(err, "could not read index.html")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = w.Write(bt)

		if err != nil {
			log.Error(err, "error writing index.html")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
