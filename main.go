package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vandit1604/go-article-manager/handlers"
	"github.com/vandit1604/go-article-manager/models"

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

	var err error

	models.DB, err = sql.Open("postgres", "postgres://postgres:mypassword@localhost/article-manager?sslmode=disable")
	if err != nil {
		log.Fatalf("error creating the DB: %v", err)
	}
}

func main() {
	// Default returns an Engine instance with the Logger and Recovery middleware already attached.
	router := gin.Default()
	// load the html files under templates
	// Once loaded, these don’t have to be read again on every request making Gin web applications very fast.
	router.LoadHTMLGlob("templates/*")

	router.GET("/", handlers.ShowIndexPage)
	router.GET("/article/view/:article_id", handlers.ShowArticle)
	router.GET("/post-article", handlers.RegisterArticle)

	router.Run()
}
