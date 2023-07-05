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

	id := c.Param("id")

	updateData := model.UserUpdate{}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "binding error"})
		return
	}

	fmt.Println(updateData.Age, updateData.Email, updateData.Name)

	rows, err := db.Query(`UPDATE * FROM users WHERE id = ?`, id)
	if err != nil {
		fmt.Println("Failed to execute the query:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database query error"})
		return
	}
	defer rows.Close()

	var users []model.UserUpdate

	for rows.Next() {
		var user model.UserUpdate
		var createdAt string

		err := rows.Scan(&user.Name, &user.Email, &user.Age, &createdAt)
		if err != nil {
			fmt.Println("Failed to scan row:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database scan error"})
			return
		}
		// parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAt)
		// if err != nil {
		// 	fmt.Println("Failed to parse created_at:", err)
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "time parsing error"})
		// 	return
		// }
		// user.CreatedAt = parsedTime
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error occurred while iterating over rows:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database row iteration error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Record Updated": body})

	// 	fmt.Println("Failed to execute the query:", err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "database query error"})
	// 	return
	// }
	// defer rows.Close()

	// var users []model.User

	// for rows.Next() {
	// 	var user model.User
	// 	var createdAt string

	// 	err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &createdAt)
	// 	if err != nil {
	// 		fmt.Println("Failed to scan row:", err)
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "database scan error"})
	// 		return
	// 	}
	// 	parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAt)
	// 	if err != nil {
	// 		fmt.Println("Failed to parse created_at:", err)
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "time parsing error"})
	// 		return
	// 	}
	// 	user.CreatedAt = parsedTime
	// 	users = append(users, user)
	// }
	// if err := rows.Err(); err != nil {
	// 	fmt.Println("Error occurred while iterating over rows:", err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "database row iteration error"})
	// 	return
	// }

	// c.JSON(http.StatusOK, users)
}
