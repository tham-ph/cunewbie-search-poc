package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/tham-ph/cunewbie-search-poc/backend-service/src/database"
	"gorm.io/gorm"
	"log"
	"time"
)

type Book struct {
	gorm.Model
	Title         string
	Author        string
	Rating        uint
	Voters        uint
	Price         float32
	Currency      string
	Description   string
	Publisher     string
	PageCount     uint
	Genres        string
	ISBN          string
	Language      string
	PublishedDate string
}

func main() {
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&Book{})
	if err != nil {
		log.Fatal(err)
	}

	//db.Create(&Book{Title: "P123", Author: "J.R.R. Tolkien", Rating: 4, Voters: 100, Price: 10.99, Currency: "USD", Description: "A great book", Publisher: "Allen & Unwin", PageCount: 295, Genres: "Fantasy", ISBN: "9780241956400", Language: "English", PublishedDate: "21 September 1937"})

	//res := Book{}
	//db.Model(&Book{}).Where("title = ?", "Poom Book").First(&res)
	//res.Title = "Poom Book 99"
	//db.Save(&res)

	db.Delete(&Book{}, 36)

	rabbitmqConnection, err := database.ConnectRabbitMQ()
	if err != nil {
		log.Fatal(err)
	}
	defer rabbitmqConnection.Close()

	rabbitmqChannel, err := rabbitmqConnection.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer rabbitmqChannel.Close()

	q, err := rabbitmqChannel.QueueDeclare("queue1", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = rabbitmqChannel.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello World again6! from queue1"),
	})
	if err != nil {
		log.Fatal(err)
	}
}
