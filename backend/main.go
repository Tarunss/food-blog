package main

import (
	"fmt"
	"net/http"

	"github.com/Tarunss/food-blog/router"
	"github.com/gorilla/handlers"
)

// defining an upgrader struct

func main() {
	r := router.Router()
	//Setting up our front end
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
	)

	// Wrap the router with CORS middleware
	wrappedRouter := cors(r)
	err := http.ListenAndServe(":8080", wrappedRouter)
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println("Hello World")
}
