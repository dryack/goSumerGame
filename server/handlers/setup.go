package handlers

import (
	"embed"
	"net/http"
)

//go:embed static
var staticFS embed.FS

func SetupHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/healthcheck", healthCheckHandler)
	mux.HandleFunc("/", indexHandler)
	//mux.HandleFunc("/api", apiHandler)

	staticFileServer := http.FileServer(http.FS(staticFS))
	mux.Handle("/static/", staticFileServer)
}
