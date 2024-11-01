package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title        string             `json:"title"        binding:"required"`
	Subtitle     string             `json:"subtitle"     binding:"required"`
	Description  string             `json:"description"  binding:"required"`
	Technologies []string           `json:"technologies" binding:"required"`
	URL          string             `json:"url"          binding:"required"`
	MonthYear    string             `json:"monthyear"    binding:"required"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
}

type ProjectResponse struct {
	ID           primitive.ObjectID `json:"id"`
	Title        string             `json:"title"`
	Subtitle     string             `json:"subtitle"`
	Description  string             `json:"description"`
	Technologies []string           `json:"technologies"`
	URL          string             `json:"url"`
	MonthYear    string             `json:"monthyear"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}
