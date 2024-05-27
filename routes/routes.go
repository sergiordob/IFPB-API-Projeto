package routes

import(
	"net/http"
	//"go.mongodb.org/mongo-driver/mongo"
	"github.com/sergiordob/IFPB-Projeto-2024/controllers"
)

func InitializeRoutes() *http.ServeMux {
	multiplexer := http.NewServeMux()

	multiplexer.HandleFunc("/api/{uf}", controllers.GetDrugstoresByState())
	multiplexer.HandleFunc("/api/{uf}/{cidade}", controllers.GetDrugstoresByCityAndState())
	
	return multiplexer
}