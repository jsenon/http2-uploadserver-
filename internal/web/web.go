package web

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
	http.ListenAndServe(":8080", nil)
}
