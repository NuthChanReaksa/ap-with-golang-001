package product

import (
	"context"
	"fmt"

	"github.com/NuthChanReaksa/ap-with-golang-001/types"
	"github.com/go-kivik/kivik"
	"github.com/google/uuid"
)

type Store struct {
	db *kivik.DB
}

func NewStore(db *kivik.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProductByID(productID string) (*types.Product, error) {
	doc := s.db.Get(context.Background(), productID)
	if doc.Err != nil { // Access the Err field directly
		return nil, fmt.Errorf("product not found: %v", doc.Err)
	}

	var product types.Product
	err := doc.ScanDoc(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (s *Store) GetProductsByID(productIDs []string) ([]types.Product, error) {
	rows, err := s.db.Find(context.Background(), map[string]interface{}{
		"selector": map[string]interface{}{
			"_id": map[string]interface{}{
				"$in": productIDs,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	var products []types.Product
	for rows.Next() {
		var product types.Product
		if err := rows.ScanDoc(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (s *Store) GetProducts() ([]*types.Product, error) {
	rows, err := s.db.AllDocs(context.Background(), kivik.Options{"include_docs": true})
	if err != nil {
		return nil, err
	}

	var products []*types.Product // Now returning slice of pointers
	for rows.Next() {
		var product types.Product
		if err := rows.ScanDoc(&product); err != nil {
			return nil, err
		}
		products = append(products, &product) // Append pointer to product
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (s *Store) CreateProduct(product types.CreateProductPayload) error {
	productDoc := map[string]interface{}{
		"_id":         generateUniqueID(),
		"name":        product.Name,
		"price":       product.Price,
		"image":       product.Image,
		"description": product.Description,
		"quantity":    product.Quantity,
	}

	_, err := s.db.Put(context.Background(), generateUniqueID(), productDoc)
	if err != nil {
		return fmt.Errorf("failed to create product: %v", err)
	}

	return nil
}

func (s *Store) UpdateProduct(product types.Product) error {
	productDoc := map[string]interface{}{
		"name":        product.Name,
		"price":       product.Price,
		"image":       product.Image,
		"description": product.Description,
		"quantity":    product.Quantity,
	}

	_, err := s.db.Put(context.Background(), product.ID, productDoc)
	if err != nil {
		return fmt.Errorf("failed to update product: %v", err)
	}

	return nil
}

func generateUniqueID() string {
	return uuid.NewString()
}
