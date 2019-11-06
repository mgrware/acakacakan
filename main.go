package main

import (
	"net/http"
	"mongo/controllers/articles"
	"github.com/gin-gonic/gin"
  "mongo/db"
  "mongo/middlewares"
)

func init() {
	db.Connect()
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.Use(middlewares.Connect)
	r.Use(middlewares.ErrorHandler)
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("articles", articles.Index)
	r.POST("articles", articles.Create)
	r.POST("articles/:_id", articles.Update)
	r.GET("articles/:_id", articles.Edit)


	return r
}

func main() {

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
