package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

type AddResult struct {
	A   int
	B   int
	Sum int
}

func add(a, b int) AddResult {
	return AddResult{a, b, a + b}
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})
	a := add(2, 2)
	router.GET("/plus", func(c *gin.Context) {
		c.JSON(http.StatusOK, a)
	})

	router.Run(":" + port)
}
