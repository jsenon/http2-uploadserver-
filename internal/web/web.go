package web

// TO DO: http2 could only be choosing if TLS is activated

import (
	"net/http"

	"github.com/jsenon/http2-uploadserver/internal/upload"
	"github.com/rs/zerolog/log"
)

// Serve luanch http server
func Serve() {
	log.Info().Msg("Startin Web Server on port 8080")
	setupRoutes()
}

func setupRoutes() {
	http.HandleFunc("/upload", upload.File)
	http.HandleFunc("/upload-ostream", upload.OStream)
	http.ListenAndServe(":8080", nil)
}
