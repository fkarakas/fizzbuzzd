package router

import (
	"github.com/fkarakas/fizzbuzzd/router/v1/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

// NewRouter creates the Gin Engine for the API server
func NewRouter(ginMode string, version string) *gin.Engine {
	gin.SetMode(ginMode)

	router := gin.Default()

	// health check
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"version": version})
		return
	})

	// FizzBuzz version 1 api group
	fizzBuzzApiV1 := router.Group("/api/v1/fizzbuzz")

	// Get Array of generated strings
	fizzBuzzApiV1.GET("/numbers/:number1/:number2/terms/:term1/:term2", routes.GetArray)

	return router
}
