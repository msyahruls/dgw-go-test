package service

import (
	"github.com/msyahruls/dgw-go-test/internal/domain"
	"github.com/msyahruls/dgw-go-test/internal/dto"
	"github.com/msyahruls/dgw-go-test/internal/repository"

	"gorm.io/gorm"
)

type ProductService interface {
	CreateProduct(req dto.CreateProductRequest) (*domain.Product, error)
	GetProducts() ([]domain.Product, error)
	GetProductByID(id uint) (*domain.Product, error)
	UpdateProduct(id uint, req dto.UpdateProductRequest) (*domain.Product, error)
	DeleteProduct(id uint) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(db *gorm.DB) ProductService {
	return &productService{
		repo: repository.NewProductRepository(db),
	}
}

func (s *productService) CreateProduct(req dto.CreateProductRequest) (*domain.Product, error) {
	product := &domain.Product{
		Name:       req.Name,
		Price:      req.Price,
		CategoryID: req.CategoryID,
	}

	if err := s.repo.Create(product); err != nil {
		return nil, err
	}
	return s.repo.FindByID(product.ID)
}

func (s *productService) GetProducts() ([]domain.Product, error) {
	return s.repo.FindAll()
}

func (s *productService) GetProductByID(id uint) (*domain.Product, error) {
	return s.repo.FindByID(id)
}

func (s *productService) UpdateProduct(id uint, req dto.UpdateProductRequest) (*domain.Product, error) {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	product.Name = req.Name
	product.Price = req.Price
	product.Category.ID = req.CategoryID

	if err := s.repo.Update(product); err != nil {
		return nil, err
	}

	return s.repo.FindByID(id)
}

func (s *productService) DeleteProduct(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
