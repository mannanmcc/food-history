package main

import (
	"github.com/gorilla/mux"
	"github.com/mannanmcc/food-history/api/handlers"
)

//NewRouter pass a request to a handler function
func NewRouter(env handlers.Env) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/item/new", env.AddNewItem).Methods("POST")
	r.HandleFunc("/item/edit/{id:[0-9]+}", env.EditItem).Methods("POST")
	r.HandleFunc("/item/list", env.GetAllItems).Methods("GET")

	return r
}
