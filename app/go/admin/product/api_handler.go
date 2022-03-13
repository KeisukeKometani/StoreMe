package product

import (
	"net/http"
	"example/online_shop/go/lib"
	"encoding/json"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func getApiHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB, p *Product, id string) {
	db.First(p, id)
}

func getAllApiHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB, products *[]Product) {
	db.Find(products)
}

func postApiHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB, p *Product) {
	lib.JsonDecodeRequestBody(r, p)
	db.Create(p)
}

func putApiHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB, p *Product, id string) {
	lib.JsonDecodeRequestBody(r, p)
	p.ID = lib.Atoui(id)
	db.Model(&p).Updates(&p)
}

func deleteApiHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB, p *Product, id string) {
	p.ID = lib.Atoui(id)
	db.Delete(&p)

}

// ApiHandler common method for base
func apiHandler(w http.ResponseWriter, r *http.Request) {
	// Get the method
	method := r.Method

	vars := mux.Vars(r)
	id := vars["id"]

	// db := database.ConnectDatabase()
	db := lib.GetDBInstance()

	var product Product
	var products []Product

	switch method {
	case "GET":
		if id == "" {
			getAllApiHandler(w, r, db, &products)
		} else {
			getApiHandler(w, r, db, &product, id)
		}
	case "POST":
		postApiHandler(w, r, db, &product)
	case "PUT":
		putApiHandler(w, r, db, &product, id)
	case "DELETE":
		deleteApiHandler(w, r, db, &product, id)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}

	// Convert product to JSON
	var b []byte
	var err error

	if id == "" && method == "GET" {
		b, err = json.Marshal(products)
	} else {
		b, err = json.Marshal(&product)
	}
	lib.ErrorCheck(err)
	
	// Write JSON to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}


// Call Handlers
func CallApiHandlers(router *mux.Router) {
	router.HandleFunc("/api/product/{id}", apiHandler).Methods("GET", "PUT", "DELETE")
	router.HandleFunc("/api/product", apiHandler).Methods("GET","POST")
}
