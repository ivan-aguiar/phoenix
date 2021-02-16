package db

import (
	"context"
	"time"
	"github.com/ivan-aguiar/phoenix/models"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

//InsertRegister es la función que inserta los datos de usuario en la DB
func InsertRegister(user models.User) (string, bool, error){

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("phoenix")
	col := db.Collection("users")

	user.Password, _ = EncryptPassword(user.Password)

	result, err := col.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}