package config

import (
	"log"
	"net/http"
	"github.com/sergiordob/IFPB-Projeto-2024/routes"
	//"go.mongodb.org/mongo-driver/mongo"
)

func InitializeServer() {
	
	server := http.Server{
		Addr:    ":8080",
		Handler: routes.InitializeRoutes(),
	}

	log.Println("Starting server at localhost:8080...")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}