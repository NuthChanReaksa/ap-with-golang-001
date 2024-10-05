package services

import (
	"github.com/google/uuid"
	models "homework1/internal/models"
	"homework1/internal/repository"
)

type ProductService interface {
	GetAllProducts() ([]models.Product, error)
	GetProductById(id uuid.UUID) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	UpdateProduct(id uuid.UUID, product models.Product) (models.Product, error)
	DeleteProduct(id uuid.UUID) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (ps *productService) GetAllProducts() ([]models.Product, error) {
	return ps.repo.GetAll()
}

func (ps *productService) GetProductById(id uuid.UUID) (models.Product, error) {
	return ps.repo.GetById(id)
}

func (ps *productService) CreateProduct(product models.Product) (models.Product, error) {
	product.ID = uuid.New()
	return ps.repo.Create(product)
}

func (ps *productService) UpdateProduct(id uuid.UUID, updatedProduct models.Product) (models.Product, error) {
	return ps.repo.Update(id, updatedProduct)
}

func (ps *productService) DeleteProduct(id uuid.UUID) error {
	return ps.repo.Delete(id)
}
