package controllers

import (
	"fmt"
	"net/http"
	"user-info/database"
	

	"github.com/gin-gonic/gin"
)

func DeleteUserById(c *gin.Context) {

	db, err := database.Connection()

	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection error"})
		return
	}
	defer db.Close()

	id := c.Param("id")
	result, err := db.Exec("DELETE from users WHERE id = ?", id)
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
		c.JSON(http.StatusAccepted, gin.H{"message": "user already deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Record Updated": "user deleted"})

}
