package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"main.go/models"
)

// UserController handles user-related requests
type UserController struct {
	DB *sqlx.DB
}

// CreateUser handles POST /users
func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert into database
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, created_at`
	err := ctrl.DB.QueryRowx(query, user.Name, user.Email).StructScan(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUsers handles GET /users
func (ctrl *UserController) GetUsers(c *gin.Context) {
	var users []models.User

	err := ctrl.DB.Select(&users, "SELECT * FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
