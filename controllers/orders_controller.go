package controllers

import (
	"net/http"
	"restAPI/database"
	"restAPI/models"

	"github.com/gin-gonic/gin"
)

// ...

func CreateOrder(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Use the DB instance from the database package
    if err := database.DB.Create(&order).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
        return
    }

    c.JSON(http.StatusCreated, order)
}

func GetOrders(c *gin.Context) {
    var orders []models.Order

    // Use the DB instance from the database package
    if err := database.DB.Find(&orders).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
        return
    }

    c.JSON(http.StatusOK, orders)
}
