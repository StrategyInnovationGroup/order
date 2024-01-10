package handler

import (
	"net/http"
	"order/pkg/models"
	"order/pkg/models/response"
	"order/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

type OrderHandler struct {
	orderService service.OrderService
}

func NewOrderHandlerImpl(service service.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: service,
	}
}

func HandleGetAllOrders(controller *gin.Context) {
	result := db.Limit(20).Find(&models.Order{})
	if result.Error != nil {
		controller.String(404, "No orders found")
	}
	controller.JSON(200, result)
}

func (handler *OrderHandler) HandleGetOrderByID(ctx *gin.Context) {
	orderId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad Request")
	}

	order := handler.orderService.FindByID(orderId)
	orderResponse := response.APIResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   order,
	}
	ctx.JSON(http.StatusOK, orderResponse)
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
