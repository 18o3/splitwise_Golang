package Handlers

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got request for adding new User\n")

	NewUserName := r.URL.Query().Get("name")
	NewUserUPI := r.URL.Query().Get("upi")

	if NewUserName == "" {
		http.Error(w, "Missing 'name' query parameter", http.StatusBadRequest)
		return
	}
	if NewUserUPI == "" {
		http.Error(w, "Missing 'upi' query parameter", http.StatusBadRequest)
		return
	}

	currentTime := time.Now()
	message := currentTime.String() + " user created"

	var newUser User
	newUser.name = NewUserName
	newUser.UserID = countID
	newUser.balance = 0
	newUser.UPI = NewUserUPI
	newUser.transactions = append(newUser.transactions, message)

	countID++
	Users = append(Users, newUser)

	io.WriteString(w, fmt.Sprintf("Added Support for a new User %s\n", NewUserName))
	io.WriteString(w, fmt.Sprintf("User is allocated the User ID %d\n", newUser.UserID))
}
