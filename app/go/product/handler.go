package product

import (
	"log"
	"net/http"
	"html/template"
	"strconv"
	"example/online_shop/database"

	"github.com/gorilla/mux"
)


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
	
	defer db.Close()

	t.Execute(w, products)
}

// EditHandler
func EditHandler(w http.ResponseWriter, r *http.Request){
	db := database.ConnectDatabase()
	// Get the id from the URL. edit/:id
	vars := mux.Vars(r)
	id := vars["id"]

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
	vars := mux.Vars(r)
	id := vars["id"]

	// When id is empty, it is a new record.
	// Insert a record.
	price, _ := strconv.Atoi(r.FormValue("price"))
	if id == "" {
		// Insert a record
		_, err := db.Exec("INSERT INTO product (name, price, description, image) VALUES (?, ?, ?, ?)", r.FormValue("name"), price, r.FormValue("description"), r.FormValue("image"))
		if err != nil {
			log.Fatal(err)
		}
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
	vars := mux.Vars(r)
	id := vars["id"]

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
	t, err := template.ParseFiles("html/product/get.html", "html/template/_head.html")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	t.Execute(w, p)
}

func stopGoRunHandler(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Stop Go Run")
}

// Call Handlers
func CallHandlers(router *mux.Router) {
	router.HandleFunc("/product/view", Viewhandler)
	router.HandleFunc("/product/edit/{id}", EditHandler)
	router.HandleFunc("/product/save/{id}", SaveHandler)
	router.HandleFunc("/product/delete/{id}", DeleteHandler)
	router.HandleFunc("/product/new", CreateHandler)
	router.HandleFunc("/product/{id}", GetHandler)
	router.HandleFunc("/product/stop", stopGoRunHandler)
}
	