package controllers

import (
	"fmt"
	"net/http"
	"user-info/database"
	"user-info/model"

	"github.com/gin-gonic/gin"
)

func UpdateUserById(c *gin.Context) {
	var body model.UserUpdate
	db, err := database.Connection()

	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection error"})
		return
	}
	defer db.Close()
	fmt.Println(body)
	id := c.Param("id")

	updateData := model.UserUpdate{}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "binding error"})
		return
	}

	fmt.Println(updateData.Age, updateData.Email, updateData.Name)

	result, err := db.Exec("UPDATE users SET name = ?, email = ?, age = ? WHERE id = ?", updateData.Name, updateData.Email, updateData.Age, id)
	if err != nil {
		fmt.Println("Failed to execute the query:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database query error"})
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Failed to retrieve the number of affected rows:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database query error"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusAccepted, gin.H{"message": "user already updated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Record Updated": updateData})

}
