package main

import (
	"log"
	"net/http"
	"os"

	"github.com/expertmaksud/go-clean-archi/entities"
	inf "github.com/expertmaksud/go-clean-archi/infrastructures"
)

func main() {

	config := inf.NewConfig()

	db, err := inf.ConnectDatabase(config)

	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
		//panic(err.Error)
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&entities.Product{})

	http.ListenAndServe(":8080", ChiRouter(db).InitRouter())
}
