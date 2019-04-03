package main

import (
	"database/sql"
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
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	router.GET("/db", dbFunc(db))

	router.Run(":" + port)
}
