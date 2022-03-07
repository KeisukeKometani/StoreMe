package product

import (
	"log"
	"net/http"
	"example/online_shop/database"
	"encoding/json"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func jsonDecode(r *http.Request, p *Product) {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&p)
	if err != nil {
		log.Fatal(err)
	}
}

func getApiHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB, p *Product, id string) {
	readRecord(db, p, id)
}

func getAllApiHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB, products *[]Product) {
	readAllRecords(db, products)
}

func postApiHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB, p *Product) {
	jsonDecode(r, p)
	createRecord(db, p)
}

func putApiHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB, p *Product, id string) {
	readRecord(db, p, id)
	jsonDecode(r, p)

	var update_product Product
	updateRecord(db, p, update_product)
}

func deleteApiHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB, p *Product, id string) {
	readRecord(db, p, id)
	deleteRecord(db, p)

}

// ApiHandler common method for base
func apiHandler(w http.ResponseWriter, r *http.Request) {
	// Get the method
	method := r.Method

	vars := mux.Vars(r)
	id := vars["id"]

	db := database.ConnectDatabase()

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
		b, err = json.Marshal(product)
	}
	if err != nil {
		log.Fatal(err)
	}
	
	// Write JSON to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}


// Call Handlers
func CallApiHandlers(router *mux.Router) {
	router.HandleFunc("/api/product", apiHandler).Methods("GET")
	router.HandleFunc("/api/product/{id}", apiHandler).Methods("GET", "PUT", "DELETE")
	router.HandleFunc("/api/product", apiHandler).Methods("POST")
}