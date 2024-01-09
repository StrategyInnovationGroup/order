package handler

import (
	"order/pkg/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func Controllers(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", handleGetAllOrders)
	routerGroup.GET("/:id", handleGetOrderByID)
	routerGroup.POST("/", handleCreateOrder)
	routerGroup.PUT("/:id", handleUpdateOrderByID)
	routerGroup.DELETE("/:id", handleDeleteOrderByID)
}

func handleGetAllOrders(controller *gin.Context) {
	result := db.Limit(20).Find(&models.Order{})
	if result.Error != nil {
		controller.String(404, "No orders found")
	}
	controller.JSON(200, result)
}

func handleGetOrderByID(controller *gin.Context) {
	controller.String(400, "Bad Request")
}

func handleCreateOrder(controller *gin.Context) {
	controller.String(400, "Bad Request")
}

func handleUpdateOrderByID(controller *gin.Context) {
	controller.String(400, "Bad Request")
}

func handleDeleteOrderByID(controller *gin.Context) {
	controller.String(400, "Bad Request")
}
