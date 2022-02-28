package main

// Do import packages for connecting to the database.
// And CRUD operations.
import (
	"net/http"
	"example/online_shop/go/product"
)

func main() {
	//Call Handler
	http.HandleFunc("/product/view", product.Viewhandler)
	http.HandleFunc("/product/new", product.CreateHandler)
	http.HandleFunc("/product/", product.GetHandler)
	http.HandleFunc("/product/edit/", product.EditHandler)
	http.HandleFunc("/product/save/", product.SaveHandler)
	http.HandleFunc("/product/delete/", product.DeleteHandler)

	http.ListenAndServe(":8080", nil)
}