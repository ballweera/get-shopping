package service

import (
	"github.com/ballweera/get-shopping/model"
)

// CustomerOrderService is service
type CustomerOrderService interface {
	AddCustomerOrder(co model.CustomerOrder)
	GetCustomerOrderGrandTotal(cartID string) float64
}
