package database

import (
	"context"
	"fmt"
	"log"
	//"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DatabaseConnection *mongo.Client
//var Context context.Context

func ConnectMongoDB() {
	//Conexao com o Banco de Dados
	
	//String de Configuracao do MondoDBAtlas
	uri := "mongodb+srv://root:root@buscafarma.xprmej7.mongodb.net/?retryWrites=true&w=majority&appName=BuscaFarma"

	
	//var cancel context.CancelFunc

	//Contexto com timout
	//Context, cancel = context.WithTimeout(context.Background(), 10*time.Second)

	//Contexto sem cancelamento
	Context := context.Background()
	//defer cancel()

	//Configura a opções do cliente
	clientOptions := options.Client().ApplyURI(uri)

	//Se conecta ao cluster
	var err error
	DatabaseConnection, err = mongo.Connect(Context, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//defer mongoClient.Disconnect(Context) //qualquer coisa, tirar isso

	//Verifica a conexão com o cluster

	//Para contexto com timeout
	//err = DatabaseConnection.Ping(context.TODO(), nil)

	//Para contexto sem timeout
	err = DatabaseConnection.Ping(Context, nil)
	if err != nil {
		log.Println("Failed to connect to MongoDB Atlas")
		fmt.Println(err)
	}

	if err == nil {
		fmt.Println("Connected to MongoDB Atlas")
	}

}
