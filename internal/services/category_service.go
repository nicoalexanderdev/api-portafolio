package services

import (
	"context"
	"errors"

	"github.com/nicoalexanderdev/api-portafolio/internal/models"
	"github.com/nicoalexanderdev/api-portafolio/internal/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// errores personalizados
var (
	ErrCategoryNotFound = errors.New("category not found")
)

// interfaz
type CategoryService interface {
	CreateCategory(ctx context.Context, category *models.Category) error
	GetAllCategories(ctx context.Context) ([]models.CategoryResponse, error)
	GetCategoryByID(ctx context.Context, id string) (*models.CategoryResponse, error)
	UpdateCategory(ctx context.Context, id string, category *models.Category) error
	DeleteCategory(ctx context.Context, id string) error
}

// struct con conecta con repository
type categoryService struct {
	repo repositories.CategoryRepository
}

// constructor
func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{
		repo: repo,
	}
}

// implementacion de la interfaz
func (s *categoryService) CreateCategory(ctx context.Context, category *models.Category) error {
	return s.repo.Create(ctx, category)
}

func (s *categoryService) GetAllCategories(ctx context.Context) ([]models.CategoryResponse, error) {
	categories, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	response := make([]models.CategoryResponse, len(categories))
	for i, category := range categories {
		response[i] = models.CategoryResponse{
			ID:          category.ID,
			URLNAME:     category.URLNAME,
			Name:        category.Name,
			Description: category.Description,
			CreatedAt:   category.CreatedAt,
			UpdatedAt:   category.UpdatedAt,
		}
	}

	return response, nil
}

func (s *categoryService) GetCategoryByID(ctx context.Context, id string) (*models.CategoryResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidID
	}

	category, err := s.repo.FindByID(ctx, objectID)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, ErrCategoryNotFound
	}

	return &models.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		URLNAME:     category.URLNAME,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}, nil
}

func (s *categoryService) UpdateCategory(ctx context.Context, id string, category *models.Category) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID
	}

	return s.repo.Update(ctx, objectID, category)
}

func (s *categoryService) DeleteCategory(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID
	}

	return s.repo.Delete(ctx, objectID)
}
