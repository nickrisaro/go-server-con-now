package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// ServeHTTP atiende pedidos y genera una respuesta
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if router == nil {
		fmt.Println("El router es nulo")
		router = gin.Default()
		router.GET("", func(c *gin.Context) {
			c.String(200, "Hola GIN")
		})
		router.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	router.ServeHTTP(w, r)
}
