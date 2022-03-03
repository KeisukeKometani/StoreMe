package product

import (
	"log"
	"net/http"
	"example/online_shop/database"
	"encoding/json"
	"database/sql"

	"github.com/gorilla/mux"
)


// Restful API Handler for GET/:id return
func GetApiHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, p *Product, id string) {
	// Read the record.
	ReadRecord(db, p, id)
}

// Restful API Handler for GET all
func GetAllApiHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, products *[]Product) {
	// Read all records.
	ReadAllRecords(db, products)
}

// Restful API Handler for POST
func PostApiHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, p *Product) {
	// Read the body
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&p)
	if err != nil {
		log.Fatal(err)
	}

	// Create the record.
	CreateRecord(db, p)

}

// Restful API Handler for PUT/:id
func PutApiHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, p *Product, id string) {
	// Read the record.
	ReadRecord(db, p, id)

	// Read the body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		log.Fatal(err)
	}

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

	// Update the record.
	UpdateRecord(db, p)

}

// Restful API Handler for DELETE/:id
func DeleteApiHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, p *Product, id string) {
	// Read the record.
	ReadRecord(db, p, id)

	// Update the DeletedAt column. time now.
	DeleteRecord(db, p)

}

// ApiHandler common method for base
func ApiHandler(w http.ResponseWriter, r *http.Request) {
	// Get the method
	method := r.Method

	vars := mux.Vars(r)
	id := vars["id"]

	// connect to database
	db := database.ConnectDatabase()

	var product Product
	var products []Product
	// Switch the method
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

	// close database
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