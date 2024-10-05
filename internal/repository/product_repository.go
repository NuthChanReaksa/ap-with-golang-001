package repository

import (
    "github.com/google/uuid"
    "gorm.io/gorm"
    "homework1/internal/models"
    "fmt"
)

type ProductRepository interface {
	GetAll() ([]model.Product, error)
	GetById(id uuid.UUID) (model.Product, error)
	Create(product model.Product) (model.Product, error)
	Update(id uuid.UUID,Product model.Product) (model.Product, error)
	Delete(id uuid.UUID) error
}


type productRepository struct {
    db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
    return &productRepository{db: db}
}

func (r *productRepository) GetAll() ([]model.Product, error) {
    var products []model.Product
    if err := r.db.Find(&products).Error; err != nil {
        return nil, err
    }

    return products, nil
}

func (r *productRepository) GetById(id uuid.UUID) (model.Product, error) {
    var p model.Product
    if err := r.db.First(&p, "id = ?", id).Error; err != nil {
        return p, err
    }
    return p, nil
}

func (r *productRepository) Create(p model.Product) (model.Product, error) {
    var user model.User
    
    // Check if the user with the given UserID exists
    if err := r.db.First(&user, "id = ?", p.UserID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return p, fmt.Errorf("user with ID %s not found", p.UserID)
        }
        return p, err
    }

    // If user exists, create the product
    if err := r.db.Create(&p).Error; err != nil {
        return p, err
    }

    return p, nil
}


func (r *productRepository) Update(id uuid.UUID, updatedProduct model.Product) (model.Product, error) {
    var existingProduct model.Product
    if err := r.db.First(&existingProduct, "id = ?", id).Error; err != nil {
        return existingProduct, err
    }

    // Update fields
    existingProduct.Name = updatedProduct.Name
    existingProduct.Price = updatedProduct.Price
    existingProduct.Quantity = updatedProduct.Quantity

    if err := r.db.Save(&existingProduct).Error; err != nil {
        return existingProduct, err
    }
    return existingProduct, nil
}

func (r *productRepository) Delete(id uuid.UUID) error {
    if err := r.db.Delete(&model.Product{}, "id = ?", id).Error; err != nil {
        return err
    }
    return nil
}