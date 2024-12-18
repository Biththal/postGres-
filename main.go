package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"main.go/routes"
)

func main() {
	// Database connection
	db, err := sqlx.Connect("postgres", "postgres://postgres:admin@localhost:5432/mvc_example?sslmode=disable")
	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize Gin
	router := gin.Default()

	// Setup routes
	routes.UserRoutes(router, db)

	// Run server
	log.Println("Server running on http://localhost:8080")
	router.Run(":8080")
}
