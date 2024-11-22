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
	ErrBlogNotFound = errors.New("blog not found")
)

type BlogService interface {
	CreateBlog(ctx context.Context, blog *models.Blog) error
	GetAllBlogs(ctx context.Context) ([]models.BlogResponse, error)
	GetBlogByID(ctx context.Context, id string) (*models.BlogResponse, error)
	UpdateBlog(ctx context.Context, id string, blog *models.Blog) error
	DeleteBlog(ctx context.Context, id string) error
}

type blogService struct {
	repo repositories.BlogRepository
}

func NewBlogService(repo repositories.BlogRepository) BlogService {
	return &blogService{
		repo: repo,
	}
}

func (s *blogService) CreateBlog(ctx context.Context, blog *models.Blog) error {
	return s.repo.Create(ctx, blog)
}

func (s *blogService) GetAllBlogs(ctx context.Context) ([]models.BlogResponse, error) {
	blogs, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	response := make([]models.BlogResponse, len(blogs))
	for i, blog := range blogs {
		response[i] = models.BlogResponse{
			ID:           blog.ID,
			Title:        blog.Title,
			Subtitle:     blog.Subtitle,
			Duration:     blog.Duration,
			Content:      blog.Content,
			ExamplePaths: blog.ExamplePaths,
			CreatedAt:    blog.CreatedAt,
			UpdatedAt:    blog.UpdatedAt,
		}
	}

	return response, nil
}

func (s *blogService) GetBlogByID(ctx context.Context, id string) (*models.BlogResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidID
	}

	blog, err := s.repo.FindByID(ctx, objectID)
	if err != nil {
		return nil, err
	}

	if blog == nil {
		return nil, ErrBlogNotFound
	}

	return &models.BlogResponse{
		ID:           blog.ID,
		Title:        blog.Title,
		Subtitle:     blog.Subtitle,
		Duration:     blog.Duration,
		Content:      blog.Content,
		ExamplePaths: blog.ExamplePaths,
		CreatedAt:    blog.CreatedAt,
		UpdatedAt:    blog.UpdatedAt,
	}, nil
}

func (s *blogService) UpdateBlog(ctx context.Context, id string, blog *models.Blog) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID
	}

	return s.repo.Update(ctx, objectID, blog)
}

func (s *blogService) DeleteBlog(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID
	}

	return s.repo.Delete(ctx, objectID)
}
