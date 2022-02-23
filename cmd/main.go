package main

import (
	"log"
	"os"

	"github.com/do-it-if-i-can/server/lib"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Setting up Gin
	r := gin.Default()
	r.POST("/query", lib.GraphqlHandler())
	r.GET("/", lib.PlaygroundHandler())
	r.Run()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
}
