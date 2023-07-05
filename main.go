package main

import (
	"fmt"
	"user-info/controllers"
	"user-info/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Connection()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	defer db.Close()

	router := gin.Default()
	router.POST("/user", controllers.PostUser)
	router.GET("/view", controllers.GetUser)
	router.GET("/view/:id", controllers.GetUserByID)
	router.POST("/view/:id", controllers.UpdateUserById)

	err = router.Run(":8080")
	if err != nil {
		fmt.Println("Failed to start the server:", err)
	}
}
