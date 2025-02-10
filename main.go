package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/NochboolPrime/graphql-blog/config"
	"github.com/NochboolPrime/graphql-blog/database"
	"github.com/NochboolPrime/graphql-blog/graph"
	"github.com/NochboolPrime/graphql-blog/graph/generated"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	database.ConnectDatabase()

	router := gin.Default()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.POST("/graphql", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

	router.GET("/", func(c *gin.Context) {
		playground.Handler("GraphQL Playground", "/graphql").ServeHTTP(c.Writer, c.Request)
	})

	router.Run(":8080")
}
