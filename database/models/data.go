package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Drugstore struct {
	ID           primitive.ObjectID `bson:"_id"`
	Nome         string             `bson:"nome"`
	LongLat      []float64         `bson:"longLat"`
	NomeFantasia string             `bson:"nomeFantasia"`
	Endereco     struct {
		Rua       string `bson:"rua"`
		Numero    string `bson:"numero"`
		Bairro    string `bson:"bairro"`
		Municipio string `bson:"municipio"`
		Estado    string `bson:"estado"`
	} `bson:"endereco"`
}