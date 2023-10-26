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

	headersOk := handlers.AllowedHeaders([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:5173"})

	err := http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r))
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println("Hello World")
}
