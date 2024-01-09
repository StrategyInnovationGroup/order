package handler

import "github.com/gin-gonic/gin"

func Controllers(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", handleGetAllOrder)
	routerGroup.GET("/:id", handleGetOrderByID)
	routerGroup.POST("/", handleCreateOrder)
	routerGroup.PUT("/:id", handleUpdateOrderByID)
	routerGroup.DELETE("/:id", handleDeleteOrderByID)
}

func handleGetAllOrder(controller *gin.Context) {
	controller.String(400, "Bad Request")
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
