package model

// Cart is shopping cart
type Cart struct {
	CartID     string
	CartItems  []CartItem
	Customer   Customer
	TotalPrice float64
}

func (cart *Cart) GetCartItems() []CartItem {
	return cart.CartItems
}
