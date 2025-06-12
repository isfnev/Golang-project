package main

import (
	"fmt"
	"net/http"

	router "github.com/isfnev/postgres_integration/Router"
)

func main() {
	r := router.Router()

	fmt.Println("Starting server at port 8080..")
	panic(http.ListenAndServe(":8080", r))
}
