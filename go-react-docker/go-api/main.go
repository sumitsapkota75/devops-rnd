package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

// User struct represents the user information
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// when receive the request, print the greeting meassage
		fmt.Fprint(w, "API is runnning....")

	})
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			{ID: 1, Name: "John Doe", Email: "john@example.com"},
			{ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
		}
		// Encode the user list to JSON format
		usersJSON, err := json.Marshal(users)
		if err != nil {
			http.Error(w, "Error encoding users to JSON", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(usersJSON)
	})
	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
