package service

import "github.com/ballweera/get-shopping/model"

import "github.com/ballweera/get-shopping/service"

// CustomerOrderServiceImpl is implementation of CustomerOrderService
type CustomerOrderServiceImpl struct{}

// AddCustomerOrder adds customer's order
func (cos *CustomerOrderServiceImpl) AddCustomerOrder(co model.CustomerOrder) {

}

// GetCustomerOrderGrandTotal calculates grand total of customer's orders
func (cos *CustomerOrderServiceImpl) GetCustomerOrderGrandTotal(cartID string) float64 {
	sv := service.CartServiceImpl{}
	cart := sv.GetCartByCartID(cartID)
	var total float64

	for item := range cart.GetCartItems() {
		total += item.Price
	}
	return total
}
