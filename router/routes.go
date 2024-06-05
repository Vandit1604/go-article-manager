package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vandit1604/go-article-manager/handlers"
)

func Run() {
	router := gin.Default()
	// load the html files under templates
	// Once loaded, these don’t have to be read again on every request making Gin web applications very fast.
	router.LoadHTMLGlob("templates/*")

	setUpRoutes(router)

	router.Run()
}

func setUpRoutes(router *gin.Engine) {
	// TODO: enable serving static files and use image on 404 page

	// 404 page
	router.NoRoute(handlers.ShowNotFoundPage)

	router.GET("/", handlers.ShowIndexPage)
	router.GET("/article/view/:article_id", handlers.ShowArticle)
	router.GET("/post-article", handlers.RegisterArticlePage)
	router.POST("/register-article", handlers.RegisterArticle)
	router.GET("/login", handlers.ShowLoginPage)
	router.POST("/login", handlers.AuthenticateLogin)
	router.GET("/signup", handlers.ShowSignUpPage)
	router.POST("/signup", handlers.AuthenticateSignUp)
}
