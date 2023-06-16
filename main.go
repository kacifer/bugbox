package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kacifer/buggy_gin"
	"github.com/surfinggo/mc"
	"net/http"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	buggy_gin.UseAll(r)

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

func main() {
	r := setupRouter()

	// Listen and Server in 0.0.0.0:8080
	mc.Must(r.Run(":8080"))
}
