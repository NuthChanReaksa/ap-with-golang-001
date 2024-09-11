package cart

import (
	"fmt"
	"github.com/sikozonpc/ecom/types"
)

func getCartItemsIDs(items []types.CartCheckoutItem) ([]string, error) {
	var ids []string
	for _, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for product ID %s", item.ProductID)
		}
		ids = append(ids, item.ProductID)
	}
	return ids, nil
}

func checkIfCartIsInStock(cartItems []types.CartCheckoutItem, products map[string]types.Product) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range cartItems {
		product, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("product %s is not available in the store, please refresh your cart", item.ProductID)
		}

		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %s is not available in the quantity requested", product.Name)
		}
	}

	return nil
}

func calculateTotalPrice(cartItems []types.CartCheckoutItem, products map[string]types.Product) float64 {
	var total float64

	for _, item := range cartItems {
		product := products[item.ProductID]
		total += product.Price * float64(item.Quantity)
	}

	return total
}
