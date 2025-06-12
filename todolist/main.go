package main

import (
	router "Todo/Router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router.Router())) 
}
