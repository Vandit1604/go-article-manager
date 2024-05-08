package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vandit1604/go-article-manager/handlers"
)

func main() {
	// Default returns an Engine instance with the Logger and Recovery middleware already attached.
	router := gin.Default()
	// load the html files under templates
	// Once loaded, these don’t have to be read again on every request making Gin web applications very fast.
	router.LoadHTMLGlob("templates/*")

	router.GET("/", handlers.ShowIndexPage)
	router.GET("/article/view/:article_id", handlers.ShowArticle)

	router.Run()
}
