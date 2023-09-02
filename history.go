package Handlers

import (
	"fmt"
	"io"
	"net/http"
)

func History(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got request for showing user History\n")

	userRequested := r.URL.Query().Get("name")

	for _, ActiveUser := range Users {

		if ActiveUser.name == userRequested || userRequested == "" {
			io.WriteString(w, fmt.Sprint(ActiveUser.name, "\n"))
			io.WriteString(w, fmt.Sprint(ActiveUser.UserID, "\n"))
			io.WriteString(w, fmt.Sprint(ActiveUser.UPI, "\n"))
			io.WriteString(w, fmt.Sprint(ActiveUser.balance, "\n"))

			io.WriteString(w, fmt.Sprint("User History\n"))
			for _, transaction := range ActiveUser.transactions {
				io.WriteString(w, fmt.Sprint(transaction, "\n"))
			}

			io.WriteString(w, "__________________________\n")
		}
	}
}
