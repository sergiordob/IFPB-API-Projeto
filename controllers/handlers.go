package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/sergiordob/IFPB-Projeto-2024/database"
	"github.com/sergiordob/IFPB-Projeto-2024/database/models"
	"go.mongodb.org/mongo-driver/bson"
)

func EndPoint01() http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {
        // Seleciona o banco de dados e a coleção
        databaseHandler := database.DatabaseConnection.Database("farma")
        collection := databaseHandler.Collection("farma_full_collection")

        // Filtra as farmacias por estado (UF) - Configurado para PB
        filter := bson.M{"endereco.estado": "PB"}

        // Encontra os documentos correspondentes
        cursor, err := collection.Find(context.TODO(), filter)
        if err != nil {
            panic(err)
        }
        defer cursor.Close(context.TODO())

        // Cria um slice de structs para armazenar as farmacias
        var drugstores []models.Drugstore

        // Itera sobre os resultados
        for cursor.Next(context.TODO()) {
            var drugstore models.Drugstore

            // Decodifica o documento
            if err := cursor.Decode(&drugstore); err != nil {
                panic(err)
            }

            // Adiciona o documento à lista
            drugstores = append(drugstores, drugstore)
        }

        // Verifica se o cursor possui erro
        if err := cursor.Err(); err != nil {
            panic(err)
        }

        // Exibe as informações das farmácias
        if len(drugstores) > 0 {
            for _, f := range drugstores {
                fmt.Fprintf(writer, "Nome: %s\nEndereço: %s, %s - %s\n\n", f.Nome, f.Endereco.Rua, f.Endereco.Numero, f.Endereco.Municipio)
            }
        } else {
            fmt.Fprintln(writer, "Nenhuma farmácia encontrada.")
        }
    }
}

func EndPoint02() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//Acessando o banco de dados 'farma' e listando as collections:

		if database.DatabaseConnection == nil {
			log.Fatal("Erro: conexão com o banco de dados não foi estabelecida")
		}
		//Estabelece conexão com o banco de dados 'farma'
		database := database.DatabaseConnection.Database("farma")

		//Printa no terminal os dados da collection 'farma_full_collection'
		collection := database.Collection("farma_full_collection")
		cursor, err := collection.Find(context.TODO(), bson.M{})
		if err != nil {
			log.Fatal(err)
		}

		var results []models.Drugstore
		if err = cursor.All(context.TODO(), &results); err != nil {
			log.Fatal(err)
		}

		//log.Println(len(results))
		for _, result := range results {
			fmt.Println("ID: ", result.ID.Hex())
			fmt.Println("Nome: ", result.Nome)
			fmt.Println("Latitude: ", fmt.Sprintf("%f", result.LongLat[0]))
			fmt.Println("Longitude: ", fmt.Sprintf("%f", result.LongLat[1]))
			fmt.Println("Nome Fantasia: ", result.NomeFantasia)
			fmt.Println("Endereço: ")
			fmt.Println("\tRua: ", result.Endereco.Rua)
			fmt.Println("\tNúmero: ", result.Endereco.Numero)
			fmt.Println("\tBairro: ", result.Endereco.Bairro)
			fmt.Println("\tMunicípio: ", result.Endereco.Municipio)
			fmt.Println("\tEstado: ", result.Endereco.Estado)
			fmt.Println("-----------------------------")
		}
	}
}
