package product

import (
	"log"
	"net/http"
	"example/online_shop/database"
	"encoding/json"
	"database/sql"

	"github.com/gorilla/mux"
)

func JsonDecode(r *http.Request, p *Product) {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&p)
	if err != nil {
		log.Fatal(err)
	}
}

func GetApiHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, p *Product, id string) {
	ReadRecord(db, p, id)
}

func GetAllApiHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, products *[]Product) {
	ReadAllRecords(db, products)
}

func PostApiHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, p *Product) {
	JsonDecode(r, p)
	CreateRecord(db, p)
}

func PutApiHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, p *Product, id string) {
	ReadRecord(db, p, id)
	JsonDecode(r, p)

	// Check column present or not
	if p.Name != "" {
		p.Name = p.Name
	}
	if p.Price != 0 {
		p.Price = p.Price
	}
	if p.Description != "" {
		p.Description = p.Description
	}
	if p.Image != "" {
		p.Image = p.Image
	}

	UpdateRecord(db, p)
}

func DeleteApiHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, p *Product, id string) {
	ReadRecord(db, p, id)
	DeleteRecord(db, p)

}

// ApiHandler common method for base
func ApiHandler(w http.ResponseWriter, r *http.Request) {
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
			GetAllApiHandler(w, r, db, &products)
		} else {
			GetApiHandler(w, r, db, &product, id)
		}
	case "POST":
		PostApiHandler(w, r, db, &product)
	case "PUT":
		PutApiHandler(w, r, db, &product, id)
	case "DELETE":
		DeleteApiHandler(w, r, db, &product, id)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}

	defer db.Close()

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
	router.HandleFunc("/api/product", ApiHandler).Methods("GET")
	router.HandleFunc("/api/product/{id}", ApiHandler).Methods("GET", "PUT", "DELETE")
	router.HandleFunc("/api/product", ApiHandler).Methods("POST")
}