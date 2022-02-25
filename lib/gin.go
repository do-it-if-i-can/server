package lib

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/do-it-if-i-can/server/graph"
	"github.com/do-it-if-i-can/server/graph/generated"
	"github.com/gin-gonic/gin"
)

// ref: https://github.com/99designs/gqlgen/blob/master/docs/content/recipes/gin.md

func GraphqlHandler() gin.HandlerFunc {
	// connect db
	db, err := ConnectDatabase()
	if err != nil {
		panic(err)
	}
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
