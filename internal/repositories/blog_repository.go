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

type BlogRepository interface {
	Create(ctx context.Context, blog *models.Blog) error
	FindAll(ctx context.Context) ([]models.Blog, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*models.Blog, error)
	Update(ctx context.Context, id primitive.ObjectID, blog *models.Blog) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}

type blogRepository struct {
	collection *mongo.Collection
}

// constructor
func NewBlogRepository(db *mongo.Database) BlogRepository {
	return &blogRepository{
		collection: db.Collection("blogs"),
	}
}

// implementacion del interface

// create
func (r *blogRepository) Create(ctx context.Context, blog *models.Blog) error {
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, blog)
	if err != nil {
		return err
	}

	blog.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

// fianall
func (r *blogRepository) FindAll(ctx context.Context) ([]models.Blog, error) {
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var blogs []models.Blog
	if err = cursor.All(ctx, &blogs); err != nil {
		return nil, err
	}

	return blogs, nil
}

// find by id
func (r *blogRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Blog, error) {
	var blog models.Blog

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&blog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &blog, nil
}

// update
func (r *blogRepository) Update(ctx context.Context, id primitive.ObjectID, blog *models.Blog) error {
	blog.UpdatedAt = time.Now()

	update := bson.M{
		"$set": bson.M{
			"title":         blog.Title,
			"subtitle":      blog.Subtitle,
			"duration":      blog.Duration,
			"content":       blog.Content,
			"example_paths": blog.ExamplePaths,
			"updated_at":    blog.UpdatedAt,
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

// delete
func (r *blogRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
