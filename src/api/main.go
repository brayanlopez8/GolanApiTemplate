package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"

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
	Router.GET("/suma/:a/:b", handlersuma)
	return Router
}
func handlersuma(c *gin.Context) {
	a, _ := strconv.Atoi(c.Param("a"))

	b, _ := strconv.Atoi(c.Param("b"))

	c.JSON(200, gin.H{
		"Resultado": a + b,
	})
}

// func validarCamposEnteros(key string) int{
// 	value, err:= strconv.Atoi(key)
// 	if(err != nil){

// 	}
// }

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
