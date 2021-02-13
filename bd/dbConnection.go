package bd
 
import (
    "context"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
 
/*MongoCN es el objeto de conexión a la base de datos*/
var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:1123581321@phoenix.shlae.mongodb.net/<dbname>?retryWrites=true&w=majority")
 
/*ConnectDB es la función que me permite conectar a la base de datos
  Devuelve una conexión a la BD del tipo Mongo Client*/
func ConnectDB() *mongo.Client{
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil{
        log.Fatal(err.Error())
        return client
    }
    err = client.Ping(context.TODO(),nil)
    if err != nil{
        log.Fatal(err.Error())
        return client
    }
    log.Println("Successful connection with database")
    return client
}
 
/*checkConnection es el ping a la BD*/
func checkConnection() int {
    err := MongoCN.Ping(context.TODO(),nil)
    if err != nil{
        return 0
    }
    return 1
}
