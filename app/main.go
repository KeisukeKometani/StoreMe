package main

// Do import packages for connecting to the database.
// And CRUD operations.
import (
	"net/http"
	"example/online_shop/go/admin/product"
	"example/online_shop/go/lib"
	"example/online_shop/go/database/connect"
	"example/online_shop/go/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	lib.SetDBConnecter(connect.DBConnect{})

	routes.SetupRoutes(router);
	//Call Handler
	product.CallHandlers(router)
	product.CallApiHandlers(router)

	fs := http.FileServer(http.Dir("assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	http.ListenAndServe(":8080", router)
}
