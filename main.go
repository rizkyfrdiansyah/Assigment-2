package main

import (
	"restAPI/controllers"
	"restAPI/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Connect to the database
	database.Connect()

	// API routes
	v1 := router.Group("/api/v1")
	// Endpoints
	{
		v1.POST("/orders", controllers.CreateOrder)
		v1.GET("/orders", controllers.GetOrders)
	}
	// router.POST("/orders", controllers.CreateOrder)
	// router.GET("/orders", controllers.GetOrders)

	// Run the server
	router.Run(":8000")
}
