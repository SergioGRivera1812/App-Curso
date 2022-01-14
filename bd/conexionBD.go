package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN es el objeto de conexion a la base de datos
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://sergio:didier@clusterappcurso.ket5t.mongodb.net/twittor?retryWrites=true&w=majority")

// Conectar BD es la que conecta la base de datos
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions) // Conexion a base de datos
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil) // Conexion a ping
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Exitosa perra!:)")
	return client
}

// ChequeoConnection es el ping a la base de datos
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
