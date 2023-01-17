package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Quote struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Quote    string             `json:"quote" bson:"quote" validate:"required"`
	Author   string             `json:"author,omitempty" bson:"author,omitempty" validate:"required"`
	Category string             `json:"category,omitempty" bson:"category,omitempty" validate:"required"`
	Tags     []string           `json:"tags,omitempty" bson:"tags,omitempty" validate:"required"`
	Likes    int                `json:"likes" bson:"likes"`
}
