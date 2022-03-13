package routes

import (
	"example/online_shop/go/lib"
	"example/online_shop/go/admin/transaction"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router){
	transaction := transaction.New()
	lib.DefineRoutes(r, transaction)
}
