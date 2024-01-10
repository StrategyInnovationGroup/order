package router

import (
	"order/pkg/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(orderController *handler.OrderHandler) *gin.Engine {
	router := gin.Default()

	baseRouterGroup := router.Group("/api/v1/")
	baseRouterGroup.GET("ping", handlePing)
	orderRouterGroup := baseRouterGroup.Group("/order")

	controllers(orderRouterGroup, orderController)

	return router
}

func controllers(routerGroup *gin.RouterGroup, orderController *handler.OrderHandler) {
	//routerGroup.GET("/", orderController.HandleGetAllOrders)
	routerGroup.GET("/:id", orderController.HandleGetOrderByID)
	//routerGroup.POST("/", orderController.handleCreateOrder)
	//routerGroup.PUT("/:id", orderController.handleUpdateOrderByID)
	//routerGroup.DELETE("/:id", orderController.handleDeleteOrderByID)
}

func handlePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
