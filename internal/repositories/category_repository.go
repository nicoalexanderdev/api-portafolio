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

type CategoryRepository interface {
	Create(ctx context.Context, category *models.Category) error
	FindAll(ctx context.Context) ([]models.Category, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*models.Category, error)
	Update(ctx context.Context, id primitive.ObjectID, category *models.Category) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}

type categoryRepository struct {
	collection *mongo.Collection
}

// constructor
func NewCategoryRepository(db *mongo.Database) CategoryRepository {
	return &categoryRepository{
		collection: db.Collection("categories"),
	}
}

func (r *categoryRepository) Create(ctx context.Context, category *models.Category) error {
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, category)
	if err != nil {
		return err
	}

	category.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *categoryRepository) FindAll(ctx context.Context) ([]models.Category, error) {
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var categories []models.Category
	if err = cursor.All(ctx, &categories); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *categoryRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Category, error) {
	var category models.Category

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&category)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepository) Update(ctx context.Context, id primitive.ObjectID, category *models.Category) error {
	category.UpdatedAt = time.Now()

	update := bson.M{
		"$set": bson.M{
			"name":        category.Name,
			"description": category.Description,
			"updated_at":  category.UpdatedAt,
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

func (r *categoryRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
