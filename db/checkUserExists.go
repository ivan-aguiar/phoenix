package db

import (
	"context"
	"time"

	"github.com/ivan-aguiar/phoenix/models"
	"go.mongodb.org/mongo-driver/bson"
)

//CheckUserExists recibe como parámetro un email y comprueba si ya existe en la DB
func CheckUserExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("phoenix")
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
