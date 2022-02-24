package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
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
	r.Use(RequestLogger()) // logging request body
	r.POST("/query", lib.GraphqlHandler())
	r.GET("/", lib.PlaygroundHandler())
	r.Run()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because rdr1 will be read.

		fmt.Println(readBody(rdr1)) // Print request body

		c.Request.Body = rdr2
		c.Next()
	}
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}
