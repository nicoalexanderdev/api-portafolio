package models

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title      string             `json:"title" binding:"required"`
	UrlName    string             `json:"urlname"`
	Subtitle   string             `json:"subtitle" binding:"required"`
	Duration   int                `json:"duration" binding:"required"`
	Content    json.RawMessage    `json:"content" binding:"required"`
	Images     []string           `json:"images"`
	CategoryId primitive.ObjectID `json:"categoryId" bson:"categoryId" binding:"required"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
}

type BlogResponse struct {
	ID         primitive.ObjectID `json:"id"`
	Title      string             `json:"title"`
	UrlName    string             `json:"urlname"`
	Subtitle   string             `json:"subtitle"`
	Duration   int                `json:"duration"`
	Content    json.RawMessage    `json:"content"`
	Images     []string           `json:"images"`
	CategoryId primitive.ObjectID `json:"categoryId"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}
