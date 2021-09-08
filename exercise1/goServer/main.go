package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	APPID := os.Getenv("APPID")
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, APPID+" /")
	})

	r.GET("/app1", func(c *gin.Context) {
		c.String(http.StatusOK, APPID+" /app1")
	})

	r.GET("/app2", func(c *gin.Context) {
		c.String(http.StatusOK, APPID+" /app2")
	})

	r.GET("/admin", func(c *gin.Context) {
		c.String(http.StatusOK, APPID+" /admin")
	})
	r.Run(":9999")
}
