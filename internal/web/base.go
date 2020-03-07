package web

import (
	"fmt"
	"net/http"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("OK\n")))
}
