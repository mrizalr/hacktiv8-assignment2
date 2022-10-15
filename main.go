package main

import (
	"test/handlers"

	"github.com/gin-gonic/gin"
	// "test/database"
)

func main() {
	// database.Connect(true)
	r := gin.Default()

	r.POST("/orders", handlers.CreateOrder)
	r.GET("/orders", handlers.GetOrders)
	r.PUT("/orders/:orderID", handlers.UpdateOrder)
	r.DELETE("/orders/:orderId", handlers.DeleteOrder)
	r.Run(":8080")
}
