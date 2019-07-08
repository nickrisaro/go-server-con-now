package handler

import (
	"fmt"
	"net/http"
)

// ServeHTTP atiende pedidos y genera una respuesta
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go on Now!</h1>")
}
