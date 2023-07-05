package controllers

import (
	"fmt"
	"net/http"
	"user-info/database"
	"user-info/model"

	"github.com/gin-gonic/gin"
)

func PostUser(c *gin.Context) {
	var body model.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "binding error 1"})
		return
	}
	fmt.Println("*", body)

	db, err := database.Connection()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection error"})
		return
	}
	defer db.Close()

	Query := `INSERT INTO users (name, email, age) VALUES (?, ?, ?)`

	
	_, err = db.Exec(Query, body.Name, body.Email, body.Age)
	if err != nil {
		fmt.Println("Query execution error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
