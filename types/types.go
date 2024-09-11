package types

import (
	"time"
)

// CouchDB typically uses _id and _rev fields for document identification and versioning
type CouchDBDocument struct {
	ID  string `json:"_id,omitempty"`  // Document ID in CouchDB
	Rev string `json:"_rev,omitempty"` // Document revision
}

type User struct {
	ID        string    `json:"_id,omitempty"`  // CouchDB uses string-based IDs
	Rev       string    `json:"_rev,omitempty"` // CouchDB revision
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type Product struct {
	CouchDBDocument           // Embedding to include _id and _rev fields
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Image           string    `json:"image"`
	Price           float64   `json:"price"`
	Quantity        int       `json:"quantity"`
	CreatedAt       time.Time `json:"createdAt"`
}

type CartCheckoutItem struct {
	ProductID string `json:"productID"` // Using string for CouchDB document IDs
	Quantity  int    `json:"quantity"`
}

type Order struct {
	ID      string             `json:"_id,omitempty"` // Document ID in CouchDB
	UserID  string             `json:"userId"`        // Changed from int to string
	Total   float64            `json:"total"`
	Status  string             `json:"status"`
	Address string             `json:"address"`
	Items   []CartCheckoutItem `json:"items"` // Changed from string to []CartCheckoutItem
}

type OrderItem struct {
	ID        string  `json:"_id,omitempty"` // Document ID in CouchDB
	OrderID   string  `json:"orderId"`
	ProductID string  `json:"productId"` // Changed from int to string
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

// UserStore defines the methods required for user operations.
type UserStore interface {
	GetUserByID(id string) (*User, error)
	CreateUser(user User) error
	GetUserByEmail(email string) (*User, error)
}

// ProductStore defines the methods required for product operations.
type ProductStore interface {
	GetProductByID(id string) (*Product, error)
	GetProducts() ([]*Product, error) // Slice of pointers
	CreateProduct(product CreateProductPayload) error
	UpdateProduct(product Product) error
	GetProductsByID(ids []string) ([]Product, error)
}

type OrderStore interface {
	CreateOrder(Order) (string, error) // Returns document ID instead of int
	CreateOrderItem(OrderItem) error
}

type CreateProductPayload struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CartCheckoutPayload struct {
	Items []CartCheckoutItem `json:"items" validate:"required"`
}
