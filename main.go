package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Items struct {
	gorm.Model
	ItemCode string `json:"item_code"`
	Description string `json:"description"`
	Quantity int `json:"quantity"`
	OrderID int `json:"order_id"`
}

type Orders struct {
	gorm.Model
	CustomerName string `json:"customer_name"`
	OrderedAt string `json:"ordered_at"`
}

func main() {
	var err error
	db, err = gorm.Open(postgres.Open("host=localhost user=postgres dbname=Inventory_Management sslmode=disable password=Swadaya05"), &gorm.Config{})
	if err != nil {
		panic("Gagal terhubung ke database")
	}

	db.AutoMigrate(&Items{}, &Orders{})

	router := gin.Default()

	// router.POST("/orders", getOrder)

	// router.PUT("/orders/:orderId", updateOrder)

	// router.DELETE("/orders/:orderId", deleteOrder)

	fmt.Println("Server berjalan pada port 8080")
	router.Run(":8080")
}

func createOrder(c *gin.Context) {

}