package service

import (
	"github.com/msyahruls/dgw-go-test/internal/domain"
	"github.com/msyahruls/dgw-go-test/internal/dto"
	"github.com/msyahruls/dgw-go-test/internal/repository"

	"gorm.io/gorm"
)

type CategoryService interface {
	CreateCategory(req dto.CreateCategoryRequest) (*dto.CategoryResponse, error)
	GetCategories() ([]dto.CategoryResponse, error)
	GetCategoryByID(id uint) (*dto.CategoryResponse, error)
	UpdateCategory(id uint, req dto.UpdateCategoryRequest) (*dto.CategoryResponse, error)
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

func (s *categoryService) CreateCategory(req dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {
	category := &domain.Category{
		Name: req.Name,
	}

	if err := s.repo.Create(category); err != nil {
		return nil, err
	}

	res := dto.ToCategoryResponse(category)
	return &res, nil
}

func (s *categoryService) GetCategories() ([]dto.CategoryResponse, error) {
	categories, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return dto.ToCategoryResponses(categories), nil
}

func (s *categoryService) GetCategoryByID(id uint) (*dto.CategoryResponse, error) {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	res := dto.ToCategoryResponse(category)
	return &res, nil
}

func (s *categoryService) UpdateCategory(id uint, req dto.UpdateCategoryRequest) (*dto.CategoryResponse, error) {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	category.Name = req.Name

	if err := s.repo.Update(category); err != nil {
		return nil, err
	}

	res := dto.ToCategoryResponse(category)
	return &res, nil
}

func (s *categoryService) DeleteCategory(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
