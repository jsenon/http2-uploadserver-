package web

// TO DO: http2 could only be choosing if TLS is activated

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jsenon/http2-uploadserver/internal/upload"
	"github.com/rs/zerolog/log"
)

// Serve luanch http server
func Serve() {
	log.Info().Msg("Startin Web Server on port 8080")
	setupRoutes()
}

func setupRoutes() {

	go func() {
		// service connections
		http.HandleFunc("/upload", upload.File)
		http.HandleFunc("/healthz", healthz)
		http.HandleFunc("/upload-ostream", upload.OStream)
		log.Info().Msg("Server Listening on port 8080")
		if err := http.ListenAndServe(":8080", nil); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutdown Server ...")
}
