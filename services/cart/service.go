package cart

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sikozonpc/ecom/services/auth"
	"github.com/sikozonpc/ecom/types"
	"github.com/sikozonpc/ecom/utils"
)

type Handler struct {
	store      types.ProductStore
	orderStore types.OrderStore
	userStore  types.UserStore
}

// NewHandler initializes the handler
func NewHandler(
	store types.ProductStore,
	orderStore types.OrderStore,
	userStore types.UserStore,
) *Handler {
	return &Handler{
		store:      store,
		orderStore: orderStore,
		userStore:  userStore,
	}
}

// RegisterRoutes defines the routes for the cart service
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", auth.WithJWTAuth(h.handleCheckout, h.userStore)).Methods(http.MethodPost)
}

// handleCheckout processes the checkout request
func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())

	var cart types.CartCheckoutPayload
	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to parse JSON: %v", err))
		return
	}

	if err := utils.Validate.Struct(cart); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("validation error: %v", validationErrors))
		return
	}

	productIds, err := utils.GetCartItemsIDs(cart.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error getting cart item IDs: %v", err))
		return
	}

	products, err := h.store.GetProductsByID(productIds)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error retrieving products: %v", err))
		return
	}

	orderID, totalPrice, err := h.createOrder(products, cart.Items, userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error creating order: %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"total_price": totalPrice,
		"order_id":    orderID,
	})
}

func (h *Handler) createOrder(products []types.Product, cartItems []types.CartCheckoutItem, userID string) (string, float64, error) {
	var totalPrice float64

	for _, item := range cartItems {
		var product *types.Product
		for _, p := range products {
			if item.ProductID == p.ID {
				product = &p
				break
			}
		}
		if product == nil {
			return "", 0, fmt.Errorf("product with ID %s not found", item.ProductID)
		}

		if product.Quantity < item.Quantity {
			return "", 0, fmt.Errorf("not enough stock for product: %s", product.Name)
		}

		totalPrice += float64(item.Quantity) * product.Price

		product.Quantity -= item.Quantity
		if err := h.store.UpdateProduct(*product); err != nil {
			return "", 0, fmt.Errorf("failed to update product stock: %v", err)
		}
	}

	orderID, err := h.orderStore.CreateOrder(types.Order{
		UserID: userID,
		Items:  cartItems,
		Total:  totalPrice,
	})
	if err != nil {
		return "", 0, fmt.Errorf("failed to create order: %v", err)
	}

	return orderID, totalPrice, nil
}
