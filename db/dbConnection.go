package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoConnection es el objeto de conexión a la DB*/
var MongoConnection = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:1123581321@phoenix.shlae.mongodb.net/phoenix?retryWrites=true&w=majority")

/*ConnectDB es la función que permite conectar a la DB
  Devuelve una conexión a la DB del tipo Mongo Client*/
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Successful connection with database")
	return client
}

/*CheckConnection es el ping a la DB*/
func CheckConnection() int {
	err := MongoConnection.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
