package main

import (
	"log"
	"net/http"

	"github.com/mannanmcc/food-history/api/handlers"
	"github.com/mannanmcc/food-history/api/models"
)

func main() {
	db, err := models.NewDB("web:web@tcp(db:3306)/foodhistory?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	env := handlers.Env{Db: db}
	router := NewRouter(env)

	//log.Fatal(http.ListenAndServeTLS(":8080", "server.crt", "server.key", router))
	log.Fatal(http.ListenAndServe(":3000", router))
}
