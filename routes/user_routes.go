package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"main.go/controllers"
)

func UserRoutes(router *gin.Engine, db *sqlx.DB) {
	userController := controllers.UserController{DB: db}

	// Define user routes
	router.POST("/users", userController.CreateUser)
	router.GET("/users", userController.GetUsers)
}
