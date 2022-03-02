package product

import (
	"log"
	"net/http"
	"example/online_shop/database"
	"encoding/json"
	"time"

	"github.com/gorilla/mux"
)

// Restful API Handler for GET/:id
func GetApiHandler(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDatabase()
	// Get ID from URL
	vars := mux.Vars(r)
	id := vars["id"]

	// Get the record.
	row := db.QueryRow("SELECT * FROM product WHERE id = ?", id)

	// Create a product struct.
	var p Product
	if err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.Image, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt); err != nil {
		log.Fatal(err)
	}
	// Convert product to JSON
	json, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
		return
	}

	// Write JSON to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

// Restful API Handler for GET all
func GetAllApiHandler(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDatabase()
	// Get all records.
	rows, err := db.Query("SELECT * FROM product")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate through the result set.
	products := []Product{}
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.Image, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt); err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("%+v\n", products)
		products = append(products, p)
	}

	// Convert products to JSON
	json, err := json.Marshal(products)
	if err != nil {
		log.Println(err)
		return
	}

	// Write JSON to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

// Restful API Handler for POST
func PostApiHandler(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDatabase()
	// Post the record. read the body
	decoder := json.NewDecoder(r.Body)
	var p Product
	err := decoder.Decode(&p)
	if err != nil {
		log.Fatal(err)
	}

	// Update the record.
	stmt, err := db.Prepare("INSERT INTO product (name, price, description, image) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	_ , err = stmt.Exec(p.Name, p.Price, p.Description, p.Image)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Convert products to JSON
	json, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
		return
	}

	// Write JSON to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

// Restful API Handler for PUT/:id
func PutApiHandler(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDatabase()
	// Get ID from URL
	vars := mux.Vars(r)
	id := vars["id"]

	// Read existing records.
	row := db.QueryRow("SELECT * FROM product WHERE id = ?", id)

	// Create a product struct.
	var p Product
	if err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.Image, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt); err != nil {
		log.Fatal(err)
	}

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
	stmt, err := db.Prepare("UPDATE product SET name = ?, price = ?, description = ?, image = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	_ , err = stmt.Exec(p.Name, p.Price, p.Description, p.Image, p.ID)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Convert product to JSON
	json, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
		return
	}

	// Write JSON to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

// Restful API Handler for DELETE/:id
func DeleteApiHandler(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDatabase()
	// Get ID from URL
	vars := mux.Vars(r)
	id := vars["id"]

	// Read existing records.
	row := db.QueryRow("SELECT * FROM product WHERE id = ?", id)

	// Create a product struct.
	var p Product
	if err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.Image, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt); err != nil {
		log.Fatal(err)
	}

	// Update the DeletedAt column. time now.
	stmt, err := db.Prepare("UPDATE product SET deleted_at = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	
	_ , err = stmt.Exec(time.Now(), p.ID)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Convert product to JSON
	p.DeletedAt = &[]time.Time{time.Now()}[0]
	json, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
		return
	}

	// Write JSON to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

// Call Handlers
func CallApiHandlers(router *mux.Router) {
	router.HandleFunc("/api/product", GetAllApiHandler).Methods("GET")
	router.HandleFunc("/api/product/{id}", GetApiHandler).Methods("GET")
	router.HandleFunc("/api/product", PostApiHandler).Methods("POST")
	router.HandleFunc("/api/product/{id}", PutApiHandler).Methods("PUT")
	router.HandleFunc("/api/product/{id}", DeleteApiHandler).Methods("DELETE")
}