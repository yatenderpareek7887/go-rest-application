package main

import (
	config "go-rest-application/src/config/swagger"
	"go-rest-application/src/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the User API!"})
	})
	config.SetupSwagger(r)
	basepath := r.Group("/api/v1")
	routes.SetupRouter(basepath)
	r.Run(":8080")
}
