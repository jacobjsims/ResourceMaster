package main

import (
	util "ResourceMaster/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Set up the server
	router := SetupRouter()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func SetupRouter() *gin.Engine {

	// Set up the router
	router := gin.Default()

	// Set up the application mode
	if util.GodotEnv("GO_ENV") != "production" && util.GodotEnv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if util.GodotEnv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	return router
}
