package main

import (
	"fmt"
	"net/http"

	"github.com/Jonathansoufer/go-reporting/handler"
)

func main() {
	handler := handler.New()

	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", handler)
}