package request

type CreateOrderRequest struct {
	ProductName   string `validate:"required, min=5 , max=55" json:"product_name"`
	OrderType     string `validate:"required, min=5 , max=55" json:"order_type"`
	OrderPrice    int    `validate:"required, gte>1 , lte<500" json:"order_price"`
	OrderQuantity int    `json:"order_quantity"`
}
