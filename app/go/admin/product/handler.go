package product

import (
	"net/http"
	"strconv"
	"example/online_shop/go/database"
	"example/online_shop/go/lib"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func viewhandler(w http.ResponseWriter, r *http.Request, db *gorm.DB, products []Product){
	db.Find(&products)
	lib.RenderCommon(w, r, products, "html/admin/product/view.html")
}

func getHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB, p Product, id string){
	db.First(&p, id)
	lib.RenderCommon(w, r, p, "html/admin/product/get.html")
}

func editHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB, p Product, id string){
	db.First(&p, id)
	lib.RenderCommon(w, r, p, "html/admin/product/edit.html")
}

func saveHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB, p Product, id string){
	price, _ := strconv.Atoi(r.FormValue("price"))
	p.Name = r.FormValue("name")
	p.Price = price
	p.Description = r.FormValue("description")
	p.Image = r.FormValue("image")

	if id == "" {
		db.Create(&p)
	} else {
		p.ID = lib.Atoui(id)
		db.Model(&p).Updates(&p)
	}

	http.Redirect(w, r, "/product/view", http.StatusFound)
}

// DeleteHandler(logic delete)
func deleteHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB, p Product, id string){
	p.ID = lib.Atoui(id)
	db.Delete(&p)
	http.Redirect(w, r, "/product/view", http.StatusFound)
}

func createHandler(w http.ResponseWriter, r *http.Request){
	lib.RenderCommon(w, r, nil, "html/admin/product/new.html")
}


// Handler common methods
func productHandler(w http.ResponseWriter, r *http.Request) {
	
	path := r.URL.Path[len("/product"):]
	method := r.Method

	vars := mux.Vars(r)
	id := vars["id"]
	
	db := database.ConnectDatabase()

	var product Product
	var products []Product
	switch path {
	case "/view":
		viewhandler(w, r, db, products)
	case "/new":
		createHandler(w, r)
	case "/"+id:
		switch method {
		case "GET":
			getHandler(w, r, db, product, id)
		case "POST":
			saveHandler(w, r, db, product, id)
		}
	case "/edit/"+id:
		editHandler(w, r, db, product, id)
	case "/delete/"+id:
		deleteHandler(w, r, db, product, id)
	case "":
		saveHandler(w, r, db, product, id)
	}

}

// Call Handlers
func CallHandlers(router *mux.Router) {
	router.HandleFunc("/product/view", productHandler)
	router.HandleFunc("/product/new", createHandler)
	router.HandleFunc("/product/edit/{id}", productHandler)
	router.HandleFunc("/product/delete/{id}", productHandler)
	router.HandleFunc("/product/{id}", productHandler).Methods("GET", "POST")
	router.HandleFunc("/product", productHandler).Methods("POST")
}
	