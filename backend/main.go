package main

import (
	"net/http"

	"github.com/Tarunss/food-blog/router"
)

// defining an upgrader struct

func main() {
	r := router.Router()
	//Setting up our front end
	http.ListenAndServe(":8080", r)
	//fmt.Println("Hello World")
}

//TODO: Right now, we are using websockets for a constant connection to the back end
// Instead, we will make buttons send API requests to our router object
