package db

import (
	"context"
	"fmt"
	"time"

	"github.com/ivan-aguiar/phoenix/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SearchProfile busca el perfil del usuario en la base de datos
func SearchProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoConnection.Database("phoenix")
	col := db.Collection("users")

	var profile models.User
	ObjID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": ObjID,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)

	profile.Password = ""
	if err != nil {
		fmt.Println("Registry not found " + err.Error())
		return profile, err
	}

	return profile, nil
}
