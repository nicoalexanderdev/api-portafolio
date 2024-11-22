package models

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title        string             `json:"title" binding:"required"`
	Subtitle     string             `json:"subtitle" binding:"required"`
	Duration     int                `json:"duration" binding:"required"`
	Content      json.RawMessage    `json:"content" binding:"required"`
	ExamplePaths []string           `json:"example_paths" binding:"required"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
}

type BlogResponse struct {
	ID           primitive.ObjectID `json:"id"`
	Title        string             `json:"title"`
	Subtitle     string             `json:"subtitle"`
	Duration     int                `json:"duration"`
	Content      json.RawMessage    `json:"content"`
	ExamplePaths []string           `json:"example_paths"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}
