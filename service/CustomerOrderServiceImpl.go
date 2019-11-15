package service

import "github.com/ballweera/get-shopping/model"

// CustomerOrderServiceImpl is implementation of CustomerOrderService
type CustomerOrderServiceImpl struct{}

// AddCustomerOrder adds customer's order
func (cos *CustomerOrderServiceImpl) AddCustomerOrder(co model.CustomerOrder) {

}

// CustomerOrderGrandTotal calculates grand total of customer's orders
func (cos *CustomerOrderServiceImpl) CustomerOrderGrandTotal(cartID string) float64 {
	sv := CartServiceImpl{}
	cart := sv.GetCartByCartID(cartID)
	var total float64

	for _, item := range cart.GetCartItems() {
		total += item.Price
	}
	return total
}
