package main

// Do import packages for connecting to the database.
// And CRUD operations.
import (
	"example/online_shop/infra/rest"
)

func main() {
	rest.Serve("8080")
}
