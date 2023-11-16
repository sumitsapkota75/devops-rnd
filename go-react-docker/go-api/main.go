package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// User struct represents the user information
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/", func(c *gin.Context) {
		// when receive the request, print the greeting meassage
		c.JSON(200, gin.H{
			"message": "Hello, welcome to your Gin server!",
		})

	})
	router.GET("/users", func(c *gin.Context) {
		users := []User{
			{ID: 1, Name: "John Doe", Email: "john@example.com"},
			{ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
			{ID: 3, Name: "Sumit Sapkota 1", Email: "sumit@example.com"},
		}
		c.JSON(http.StatusOK, users)
	})
	err := router.Run(":8080")
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
