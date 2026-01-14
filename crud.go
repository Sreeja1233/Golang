package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var users = []user{
	{1, "sree"},
	{2, "ja"},
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
}
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

}
func createUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newuser user
	json.NewDecoder(r.Body).Decode(&newuser)
	users = append(users, newuser)
	json.NewEncoder(w).Encode(users)

}
func updateUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}
	var updatename user
	err := json.NewDecoder(r.Body).Decode(&updatename)
	if err != nil {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	for i, u := range users {
		if u.Name == name {
			users[i].Name = updatename.Name

			json.NewEncoder(w).Encode(users[i])
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
	for i, u := range users {
		if u.Id == id {
			users = append(users[:i], users[i+1:]...)
			json.NewEncoder(w).Encode(users)
			return
		}
	}
}
func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/create", createUsers)
	http.HandleFunc("/update", updateUsers)
	http.HandleFunc("/delete", deleteUser)
	fmt.Println("server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
