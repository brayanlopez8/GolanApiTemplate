package main

import (
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// //Router Variable de inicializacion de servidor gin
// var Router *gin.Engine

//TestingMode ...
var TestingMode bool

// func init() {
// 	Router = gin.Default()
// }

func main() {
	Router := MapRoutes()
	StartServer(Router, TestingMode)
}

//MapRoutes ...
func MapRoutes() *gin.Engine {
	Router := gin.Default()
	Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return Router
}

//StartServer ...
func StartServer(engine *gin.Engine, TestingMode bool) {
	if TestingMode {
		servertest := httptest.NewServer(engine)
		log.Print(servertest)
		return
	}
	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}
	srv.ListenAndServe()
}
