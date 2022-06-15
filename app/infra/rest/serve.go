package rest

// Do import packages for connecting to the database.
// And CRUD operations.
import (
	"net/http"
)

func Serve(port string) {
	fs := http.FileServer(http.Dir("assets/"))
	router := getRouter()
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	http.ListenAndServe(":"+port, router)
}
