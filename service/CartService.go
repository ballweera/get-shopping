package service

// CartService is service interface
type CartService interface {
	GetCartByCartID(cartID string) string
}
