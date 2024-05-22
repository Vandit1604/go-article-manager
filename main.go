package main

import (
	"log"

	"github.com/vandit1604/go-article-manager/models"
	"github.com/vandit1604/go-article-manager/router"

	_ "github.com/lib/pq"
)

func init() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = ""
		dbname   = "article-manager"
	)

	err := models.InitDB("postgres://postgres:mypassword@localhost/article-manager?sslmode=disable")
	if err != nil {
		log.Fatalf("error creating the DB: %v", err)
	}
}

func main() {
	router.Run()
}
