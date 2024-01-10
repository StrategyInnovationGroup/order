package service

import (
	"order/pkg/models/response"
	"order/pkg/repository"

	"github.com/go-playground/validator/v10"
)

type OrderServiceImpl struct {
	OrderRepository repository.OrderRepository
	validate        *validator.Validate
}

// FindByID implements OrderService.
func (o *OrderServiceImpl) FindByID(orderId int) response.OrderResponse {
	result, err := o.OrderRepository.FindByID(orderId)
	if err != nil {
		panic(err)
	}

	return response.OrderResponse(result)

}

func NewOrderServiceImpl(orderRepo repository.OrderRepository, validate *validator.Validate) OrderService {
	return &OrderServiceImpl{
		OrderRepository: orderRepo,
		validate:        validate,
	}
}
