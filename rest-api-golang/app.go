package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Fone  string `json:"fone"`
}

var users = []User{
	{ID: 1, Name: "John Doe", Email: "chulipa@hotmail.com", Fone: "00-0000-0000"},
	{ID: 2, Name: "Jane Doe", Email: "cupincha@hotmail.com", Fone: "11-1234-5555"},
}

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			data, _ := json.Marshal(users)
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		case http.MethodPost:
			var user User
			json.NewDecoder(r.Body).Decode(&user)
			users = append(users, user)
			w.WriteHeader(http.StatusCreated)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Serving on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
