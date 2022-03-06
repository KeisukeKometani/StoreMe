package product

import (
	"log"
	"net/http"
	"html/template"
	"strconv"
	"database/sql"
	"example/online_shop/database"
	"example/online_shop/go/query"

	"github.com/gorilla/mux"
)

func renderCommon(w http.ResponseWriter, r *http.Request, data interface{}, fileNames ...string) {
	t, err := template.ParseFiles(fileNames...)
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, data)
}

func viewhandler(w http.ResponseWriter, r *http.Request, db *sql.DB, products []Product){
	query.ReadAllRecords(db, products, "product")
	log.Println("viewhandler: %p",products)
	renderCommon(w, r, products, "html/admin/product/view.html")
}

func getHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, p Product, id string){
	query.ReadRecord(db, &p, id, "product")
	renderCommon(w, r, p, "html/admin/product/get.html", "html/template/_head.html")
}

func editHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, p Product, id string){
	query.ReadRecord(db, &p, id, "product")
	renderCommon(w, r, p, "html/admin/product/edit.html")
}

func saveHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, p *Product, id string){
	price, _ := strconv.Atoi(r.FormValue("price"))
	p.Name = r.FormValue("name")
	p.Price = price
	p.Description = r.FormValue("description")
	p.Image = r.FormValue("image")

	if id == "" {
		createRecord(db, p)
	} else {
		p.ID, _ = strconv.Atoi(id)
		updateRecord(db, p)
	}

	http.Redirect(w, r, "/product/view", http.StatusFound)
}

// DeleteHandler(logic delete)
func deleteHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, p Product, id string){
	query.DeleteRecord(db, id, "product")
	http.Redirect(w, r, "/product/view", http.StatusFound)
}

func createHandler(w http.ResponseWriter, r *http.Request){
	renderCommon(w, r, nil, "html/admin/product/new.html")
}


func stopGoRunHandler(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Stop Go Run")
}

// Handler common methods
func productHandler(w http.ResponseWriter, r *http.Request) {
	
	path := r.URL.Path[len("/product/"):]

	vars := mux.Vars(r)
	id := vars["id"]
	
	db := database.ConnectDatabase()

	var product Product
	var products []Product
	switch path {
	case "view":
		viewhandler(w, r, db, products)
	case id:
		getHandler(w, r, db, product, id)
	case "edit/"+id:
		editHandler(w, r, db, product, id)
	case "save/"+id:
		saveHandler(w, r, db, &product, id)
	case "save/":
		saveHandler(w, r, db, &product, id)
	case "delete/"+id:
		deleteHandler(w, r, db, product, id)
	case "new":
		createHandler(w, r)
	}

	defer db.Close()
}

// Call Handlers
func CallHandlers(router *mux.Router) {
	router.HandleFunc("/product/view", productHandler)
	router.HandleFunc("/product/edit/{id}", productHandler)
	router.HandleFunc("/product/save/{id}", productHandler)
	router.HandleFunc("/product/save/", productHandler)
	router.HandleFunc("/product/delete/{id}", productHandler)
	router.HandleFunc("/product/new", createHandler)
	router.HandleFunc("/product/stop", stopGoRunHandler)
	router.HandleFunc("/product/{id}", productHandler) // 最後にしておかないと、他のpathが取れない。
}
	