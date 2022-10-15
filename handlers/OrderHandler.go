package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"test/database"
	"test/models"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	orderInput := models.Order{}

	if err := c.ShouldBindJSON(&orderInput); err != nil {
		panic(err.Error())
	}

	db, err := database.Connect(false)
	if err != nil {
		panic(err.Error())
	}
	defer database.CloseDB()

	if err := db.Create(&orderInput).Error; err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success inserting new data into database",
		"data":    orderInput,
	})
}

func GetOrders(c *gin.Context) {
	db, err := database.Connect(false)
	if err != nil {
		panic(err.Error())
	}
	defer database.CloseDB()

	orders := []models.Order{}
	err = db.Preload("Items").Find(&orders).Error
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get orders data",
		"data":    orders,
	})
}

func UpdateOrder(c *gin.Context) {
	orderID := c.Param("orderID")

	db, err := database.Connect(false)
	if err != nil {
		panic(err.Error())
	}
	defer database.CloseDB()

	payload := models.Order{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		panic(err.Error())
	}

	order := models.Order{}
	err = db.Model(&order).Where("id = ?", orderID).Updates(models.Order{
		CustomerName: payload.CustomerName,
		OrderedAt:    payload.OrderedAt,
	}).Error

	if err != nil {
		panic(err.Error())
	}

	for _, item := range payload.Items {
		newItem := models.Item{}
		if item.ID == 0 {
			intOrderID, _ := strconv.Atoi(orderID)

			newItem.ItemCode = item.ItemCode
			newItem.Description = item.Description
			newItem.Quantity = item.Quantity
			newItem.OrderID = uint(intOrderID)

			err := db.Create(&newItem).Error

			if err != nil {
				panic(err.Error())
			}

			continue
		}

		err = db.Model(&newItem).Where("order_id = ?", orderID).Updates(models.Item{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
		}).Error

		if err != nil {
			panic(err.Error())
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("success update data with id %v", orderID),
		"data":    payload,
	})
}

func DeleteOrder(c *gin.Context) {
	orderId := c.Param("orderId")

	db, err := database.Connect(false)
	if err != nil {
		panic(err.Error())
	}
	defer database.CloseDB()

	item := models.Item{}
	db.Delete(&item, "order_id = ?", orderId)

	order := models.Order{}
	db.Delete(&order, "id = ?", orderId)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("data with id %v successfully deleted", orderId),
	})
}
