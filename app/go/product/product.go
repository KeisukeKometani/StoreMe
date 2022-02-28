package product

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"time"
	"strconv"
	"example/online_shop/database"
	//"encoding/json"
)

//make type product struct
type Product struct {
	ID		  int  `json:"id" `
	Name	  string  `json:"name" `
	Price	  int  `json:"price" `
	Description string  `json:"description" `
	Image	  string  `json:"image" `
	CreatedAt time.Time  `json:"created_at" `
	UpdatedAt time.Time  `json:"updated_at" `
	DeletedAt *time.Time  `json:"deleted_at" `
}

// Insert a record into the database.
func insertRecord(p Product) {
	db := database.ConnectDatabase()
	// Insert a record.
	fmt.Printf("%+v\n", p)
	_, err := db.Exec("INSERT INTO product (name, price, description, image) VALUES (?, ?, ?, ?)", p.Name, p.Price, p.Description, p.Image)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

// Viewhandler
func Viewhandler(w http.ResponseWriter, r *http.Request){
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

	// Render the view.
	t, err := template.ParseFiles("html/product/view.html")
	if err != nil {
		log.Fatal(err)
	}

	// When passing in json format
	//bytes, err := json.Marshal(products)
	//if err != nil {
	//	log.Fatal(err)
	//}

	defer db.Close()

	t.Execute(w, products)
}

// EditHandler
func EditHandler(w http.ResponseWriter, r *http.Request){
	db := database.ConnectDatabase()
	// Get the id from the URL. edit/:id
	id := r.URL.Path[len("/product/edit/"):]

	// Get the record.
	row := db.QueryRow("SELECT * FROM product WHERE id = ?", id)

	// Create a product struct.
	var p Product
	if err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.Image, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt); err != nil {
		log.Fatal(err)
	}

	// Render the view.
	t, err := template.ParseFiles("html/product/edit.html")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	t.Execute(w, p)
}

// SaveHandler
func SaveHandler(w http.ResponseWriter, r *http.Request){
	db := database.ConnectDatabase()
	// Get the id from the URL. edit/:id
	id := r.URL.Path[len("/product/save/"):]

	// When id is empty, it is a new record.
	// Insert a record.
	price, _ := strconv.Atoi(r.FormValue("price"))
	if id == "" {
		p := Product{
			Name: r.FormValue("name"),
			Price: price,
			Description: r.FormValue("description"),
			Image: r.FormValue("image"),
		}
		insertRecord(p)
	} else {
		// Update a record.
		_, err := db.Exec("UPDATE product SET name = ?, price = ?, description = ?, image = ? WHERE id = ?", r.FormValue("name"), r.FormValue("price"), r.FormValue("description"), r.FormValue("image"), id)
		if err != nil {
			log.Fatal(err)
		}
	}

	defer db.Close()
	// Redirect to the view.
	http.Redirect(w, r, "/product/view", http.StatusFound)
}

// DeleteHandler(logic delete)
func DeleteHandler(w http.ResponseWriter, r *http.Request){
	db := database.ConnectDatabase()
	// Get the id from the URL. delete/:id
	id := r.URL.Path[len("/product/delete/"):]

	// Get the record.
	row := db.QueryRow("SELECT * FROM product WHERE id = ?", id)

	// Create a product struct.
	var p Product
	if err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.Image, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt); err != nil {
		log.Fatal(err)
	}

	// Update the record.
	_, err := db.Exec("UPDATE product SET deleted_at = NOW() WHERE id = ?", p.ID)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	// Redirect to the view.
	http.Redirect(w, r, "/product/view", http.StatusFound)
}

// CreateHandler
func CreateHandler(w http.ResponseWriter, r *http.Request){
	// Render the view.
	t, err := template.ParseFiles("html/product/new.html")
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, nil)
}

// GetHandler
func GetHandler(w http.ResponseWriter, r *http.Request){
	db := database.ConnectDatabase()
	// Get the id from the URL. get/:id
	id := r.URL.Path[len("/product/"):]

	// Get the record.
	row := db.QueryRow("SELECT * FROM product WHERE id = ?", id)

	// Create a product struct.
	var p Product
	if err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.Image, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt); err != nil {
		log.Fatal(err)
	}

	// Render the view.
	t, err := template.ParseFiles("html/product/get.html")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	t.Execute(w, p)
}