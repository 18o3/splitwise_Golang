package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"splitwise/Handlers"
)

func main() {

	Handlers.Init()

	http.HandleFunc("/", Handlers.GetRoot)
	http.HandleFunc("/newUser", Handlers.NewUser)
	http.HandleFunc("/history", Handlers.History)

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
