package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ServeHTTP atiende pedidos y genera una respuesta
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := gin.Default()
	router.GET("", func(c *gin.Context) {
		c.String(200, "Hola GIN")
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.ServeHTTP(w, r)
	fmt.Fprintf(w, "<h1>Hello from Go on Now!</h1>")
}
