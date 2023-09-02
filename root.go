package Handlers

import (
	"fmt"
	"io"
	"net/http"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Root Command\n")
	io.WriteString(w, "Welcome to SplitWise!\n")
}
