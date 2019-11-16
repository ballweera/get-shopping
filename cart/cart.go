package cart

import "github.com/ballweera/get-shopping/customer"

type Cart struct {
	CartID     string
	CartItems  []CartItem
	Customer   customer.Customer
	TotalPrice float64
}

type CartItem struct {
	Price float64
}

// ByID finds cart by ID
func ByID(ID string) Cart {
	return Cart{}
}
