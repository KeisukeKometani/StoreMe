package rest

import (
  "github.com/gorilla/mux"
)

var router *mux.Router

func init() {
  r := mux.NewRouter()
  router = r
}

func getRouter() *mux.Router {
  return router
}
