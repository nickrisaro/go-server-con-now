package handler

import (
	"fmt"
	"net/http"
)

// Handler atiende pedidos y genera una respuesta
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go on Now!</h1>")
}
