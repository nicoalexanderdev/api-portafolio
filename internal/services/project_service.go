package services

import (
	"context"
	"errors"

	"github.com/nicoalexanderdev/api-portafolio/internal/models"
	"github.com/nicoalexanderdev/api-portafolio/internal/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrProjectNotFound = errors.New("project not found")
	ErrInvalidID       = errors.New("invalid project ID")
)

type ProjectService interface {
	CreateProject(ctx context.Context, project *models.Project) error
	GetAllProjects(ctx context.Context) ([]models.ProjectResponse, error)
	GetProjectByID(ctx context.Context, id string) (*models.ProjectResponse, error)
	UpdateProject(ctx context.Context, id string, project *models.Project) error
	DeleteProject(ctx context.Context, id string) error
}

type projectService struct {
	repo repositories.ProjectRepository
}

func NewProjectService(repo repositories.ProjectRepository) ProjectService {
	return &projectService{
		repo: repo,
	}
}

func (s *projectService) CreateProject(ctx context.Context, project *models.Project) error {
	return s.repo.Create(ctx, project)
}

func (s *projectService) GetAllProjects(ctx context.Context) ([]models.ProjectResponse, error) {
	projects, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	response := make([]models.ProjectResponse, len(projects))
	for i, project := range projects {
		response[i] = models.ProjectResponse{
			ID:           project.ID,
			Title:        project.Title,
			Subtitle:     project.Subtitle,
			Description:  project.Description,
			Technologies: project.Technologies,
			URL:          project.URL,
			MonthYear:    project.MonthYear,
			CreatedAt:    project.CreatedAt,
			UpdatedAt:    project.UpdatedAt,
		}
	}

	return response, nil
}

func (s *projectService) GetProjectByID(ctx context.Context, id string) (*models.ProjectResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidID
	}

	project, err := s.repo.FindByID(ctx, objectID)
	if err != nil {
		return nil, err
	}
	if project == nil {
		return nil, ErrProjectNotFound
	}

	return &models.ProjectResponse{
		ID:           project.ID,
		Title:        project.Title,
		Subtitle:     project.Subtitle,
		Description:  project.Description,
		Technologies: project.Technologies,
		URL:          project.URL,
		MonthYear:    project.MonthYear,
		CreatedAt:    project.CreatedAt,
		UpdatedAt:    project.UpdatedAt,
	}, nil
}

func (s *projectService) UpdateProject(ctx context.Context, id string, project *models.Project) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID
	}

	return s.repo.Update(ctx, objectID, project)
}

func (s *projectService) DeleteProject(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID
	}

	return s.repo.Delete(ctx, objectID)
}
