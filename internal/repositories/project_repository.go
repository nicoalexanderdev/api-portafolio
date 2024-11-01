package repositories

import (
	"context"
	"time"

	"github.com/nicoalexanderdev/api-portafolio/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProjectRepository interface {
	Create(ctx context.Context, project *models.Project) error
	FindAll(ctx context.Context) ([]models.Project, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*models.Project, error)
	Update(ctx context.Context, id primitive.ObjectID, project *models.Project) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}

type projectRepository struct {
	collection *mongo.Collection
}

func NewProjectRepository(db *mongo.Database) ProjectRepository {
	return &projectRepository{
		collection: db.Collection("projects"),
	}
}

func (r *projectRepository) Create(ctx context.Context, project *models.Project) error {
	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, project)
	if err != nil {
		return err
	}

	project.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *projectRepository) FindAll(ctx context.Context) ([]models.Project, error) {
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var projects []models.Project
	if err = cursor.All(ctx, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *projectRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Project, error) {
	var project models.Project

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&project)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &project, nil
}

func (r *projectRepository) Update(ctx context.Context, id primitive.ObjectID, project *models.Project) error {
	project.UpdatedAt = time.Now()

	update := bson.M{
		"$set": bson.M{
			"title":        project.Title,
			"subtitle":     project.Subtitle,
			"description":  project.Description,
			"technologies": project.Technologies,
			"url":          project.URL,
			"monthyear":    project.MonthYear,
			"updated_at":   project.UpdatedAt,
		},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (r *projectRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
