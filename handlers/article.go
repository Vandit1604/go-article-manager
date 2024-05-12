package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vandit1604/go-article-manager/models"
)

func ShowIndexPage(c *gin.Context) {
	articles, err := models.GetAllArticles()
	if err != nil {
		log.Fatalf("Error get all articles: %v", err)
	}

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title":   "Article Manager",
			"payload": articles,
		},
	)
}

func ShowArticle(c *gin.Context) {
	param := c.Param("article_id")
	strParam, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		log.Fatal("Error converting the param to int")
	}
	article, err := models.GetArticle(strParam)
	if err != nil {
		c.HTML(http.StatusNotFound, "article.html", gin.H{
			"payload": article.Title,
		})
	} else {
		c.HTML(http.StatusOK, "article.html", gin.H{
			"payload": article,
			"next":    article.ID + 1,
		})
	}
}

func RegisterArticle(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		log.Fatalf("Error parsing the form: %v", err)
	}
	// get the values from the form and populate them in the db
}
