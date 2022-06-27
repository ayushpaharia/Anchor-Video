package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Video struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	VideoId      string             `json:"videoId,omitempty" bson:"videoId,omitempty"`
	Title        string             `json:"title,omitempty" bson:"title,omitempty"`
	Description  string             `json:"description,omitempty" bson:"description,omitempty"`
	PublishedAt  string             `json:"publishedAt,omitempty" bson:"publishedAt,omitempty"`
	ThumbnailURL string             `json:"thumbnailUrl,omitempty" bson:"thumbnailUrl,omitempty"`
}

type Videos []Video
