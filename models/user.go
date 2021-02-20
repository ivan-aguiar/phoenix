package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User es la estructura que se utilizará como modelo de usuario
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name,omitempty"`
	Lastname  string             `bson:"lastname" json:"lastname,omitempty"`
	Username  string             `bson:"username" json:"username,omitempty"`
	Birthday  time.Time          `bson:"birthday" json:"birthday,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Bio       string             `bson:"bio" json:"bio,omitempty"`
	Ubication string             `bson:"ubication" json:"ubication,omitempty"`
	Website   string             `bson:"website" json:"website,omitempty"`
}
