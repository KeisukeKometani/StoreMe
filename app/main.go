package main

// Do import packages for connecting to the database.
// And CRUD operations.
import (
	"net/http"
	"example/online_shop/go/product"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//Call Handler
	product.CallHandlers(router)
	product.CallApiHandlers(router)

	http.ListenAndServe(":8080", router)
}