package service

import (
	"github.com/msyahruls/dgw-go-test/internal/domain"
	"github.com/msyahruls/dgw-go-test/internal/dto"
	"github.com/msyahruls/dgw-go-test/internal/repository"

	"gorm.io/gorm"
)

type CategoryService interface {
	CreateCategory(req dto.CreateCategoryRequest) (*domain.Category, error)
	GetCategories() ([]domain.Category, error)
	GetCategoryByID(id uint) (*domain.Category, error)
	UpdateCategory(id uint, req dto.UpdateCategoryRequest) (*domain.Category, error)
	DeleteCategory(id uint) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(db *gorm.DB) CategoryService {
	return &categoryService{
		repo: repository.NewCategoryRepository(db),
	}
}

func (s *categoryService) CreateCategory(req dto.CreateCategoryRequest) (*domain.Category, error) {
	category := &domain.Category{
		Name: req.Name,
	}

	if err := s.repo.Create(category); err != nil {
		return nil, err
	}
	return category, nil
}

func (s *categoryService) GetCategories() ([]domain.Category, error) {
	return s.repo.FindAll()
}

func (s *categoryService) GetCategoryByID(id uint) (*domain.Category, error) {
	return s.repo.FindByID(id)
}

func (s *categoryService) UpdateCategory(id uint, req dto.UpdateCategoryRequest) (*domain.Category, error) {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	category.Name = req.Name

	if err := s.repo.Update(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) DeleteCategory(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
