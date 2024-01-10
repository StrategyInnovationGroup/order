package handler

import (
	"net/http"
	"order/pkg/models/response"
	"order/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService service.OrderService
}

func NewOrderHandlerImpl(service service.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: service,
	}
}

func (handler *OrderHandler) HandleGetAllOrders(ctx *gin.Context) {

	orders := handler.orderService.FindAll()

	orderResponse := response.APIResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   orders,
	}
	ctx.JSON(http.StatusOK, orderResponse)
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

func (handler *OrderHandler) HandleCreateOrder(controller *gin.Context) {
	controller.String(400, "Bad Request")
}

func (handler *OrderHandler) HandleUpdateOrderByID(controller *gin.Context) {
	controller.String(400, "Bad Request")
}

func (handler *OrderHandler) HandleDeleteOrderByID(controller *gin.Context) {
	controller.String(400, "Bad Request")
}
