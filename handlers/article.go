package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vandit1604/go-article-manager/models"
)

func ShowIndexPage(c *gin.Context) {
	articles := models.GetAllArticles()

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)
}

func ShowArticle(c *gin.Context) {
	param := c.Param("article_id")
	strParam, err := strconv.Atoi(param)
	if err != nil {
		log.Fatal("Error converting the param to int")
	}
	article, log := models.GetArticle(strParam)
	if log == "" {
		fmt.Println(article)
		c.HTML(http.StatusOK, "article.html", gin.H{
			"payload": article,
		})
	} else {
		c.JSON(http.StatusNotFound, log)
	}
}
