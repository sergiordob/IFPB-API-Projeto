package routes

import(
	"net/http"
	//"go.mongodb.org/mongo-driver/mongo"
	"github.com/sergiordob/IFPB-Projeto-2024/controllers"
)

func InitializeRoutes() *http.ServeMux {
	multiplexer := http.NewServeMux()

	multiplexer.HandleFunc("/drugstores/", controllers.EndPoint01())
	multiplexer.HandleFunc("/teste/", controllers.EndPoint02())
	
	return multiplexer
}