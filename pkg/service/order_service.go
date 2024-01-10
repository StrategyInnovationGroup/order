package service

import "order/pkg/models/response"

type OrderService interface {
	FindByID(orderId int) response.OrderResponse
}
