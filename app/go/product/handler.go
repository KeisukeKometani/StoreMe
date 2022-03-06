package product

import (
	"log"
	"net/http"
	"html/template"
	"strconv"
	"example/online_shop/database"
	"database/sql"

	"github.com/gorilla/mux"
)

func RenderCommon(w http.ResponseWriter, r *http.Request, data interface{}, fileNames ...string) {
	t, err := template.ParseFiles(fileNames...)
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, data)
}

func Viewhandler(w http.ResponseWriter, r *http.Request, db *sql.DB, products *[]Product){
	ReadAllRecords(db, products)
	RenderCommon(w, r, products, "html/product/view.html")
}

func GetHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, p *Product, id string){
	ReadRecord(db, p, id)
	RenderCommon(w, r, p, "html/product/get.html", "html/template/_head.html")
}

func EditHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, p *Product, id string){
	ReadRecord(db, p, id)
	RenderCommon(w, r, p, "html/product/edit.html")
}

func SaveHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, p *Product, id string){
	price, _ := strconv.Atoi(r.FormValue("price"))
	p.Name = r.FormValue("name")
	p.Price = price
	p.Description = r.FormValue("description")
	p.Image = r.FormValue("image")

	if id == "" {
		CreateRecord(db, p)
	} else {
		p.ID, _ = strconv.Atoi(id)
		UpdateRecord(db, p)
	}

	http.Redirect(w, r, "/product/view", http.StatusFound)
}

// DeleteHandler(logic delete)
func DeleteHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, p *Product, id string){
	ReadRecord(db, p, id)
	DeleteRecord(db, p)

	http.Redirect(w, r, "/product/view", http.StatusFound)
}

func CreateHandler(w http.ResponseWriter, r *http.Request){
	RenderCommon(w, r, nil, "html/product/new.html")
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
		Viewhandler(w, r, db, &products)
	case "edit/"+id:
		EditHandler(w, r, db, &product, id)
	case "save/"+id:
		SaveHandler(w, r, db, &product, id)
	case "save/":
		SaveHandler(w, r, db, &product, id)
	case "delete/"+id:
		DeleteHandler(w, r, db, &product, id)
	case "new":
		CreateHandler(w, r)
	case id:
		GetHandler(w, r, db, &product, id)
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
	router.HandleFunc("/product/new", CreateHandler)
	router.HandleFunc("/product/stop", stopGoRunHandler)
	router.HandleFunc("/product/{id}", productHandler) // 最後にしておかないと、他のpathが取れない。
}
	