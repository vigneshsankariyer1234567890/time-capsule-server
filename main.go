package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// Define routes here
	// router.GET("/", getHandler)
	// router.POST("/", postHandler)

	router.Run(":8080") // Run the server
}
