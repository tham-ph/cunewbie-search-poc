package main

import (
	"github.com/tham-ph/cunewbie-search-poc/backend-service/src/database"
	"log"
)

type Test_user struct {
	ID    int
	Name  string
	Email string
}

func main() {
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Test_user{})

	db.Create(&Test_user{ID: 1, Name: "Tham", Email: "Pham"})

	log.Println("Done")
}
