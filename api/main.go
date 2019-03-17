package main

import (
	"log"
	"net/http"

	"github.com/mannanmcc/food-history/api/handlers"
	"github.com/mannanmcc/food-history/api/models"
)

const (
	dbHost     = "192.168.33.10"
	dbUser     = "user"
	dbPassword = "password"
	dbName     = "prepaid-card"
)

func main() {
	db, err := models.NewDB(dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":3306)/" + dbName + "?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	env := handlers.Env{Db: db}
	router := NewRouter(env)

	//log.Fatal(http.ListenAndServeTLS(":8080", "server.crt", "server.key", router))
	log.Fatal(http.ListenAndServe(":8000", router))
}