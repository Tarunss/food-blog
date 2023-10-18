package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Defining an Upgrader Struct
// Lets us define a Read and Write Buffer Size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// We can also check the origin of our connection with CheckOrigin, which allows us to
	// Make connections from our react server
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Define a Reader function that will listen for connections on that particular web socket
func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		//write message out in our main
		fmt.Println(string(p))
		//Write message out in web socket
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

// Define our web socket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	//listen indefinitely for new messages coming in
	//through our web socket connection
	reader(ws)
}
func setUpRoutes() {
	//Sets up default route , or (/)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "SimpleServer")
	})
	http.HandleFunc("/ws", serveWs)
}
func main() {
	setUpRoutes()
	http.ListenAndServe(":8080", nil)
	fmt.Println("Hello World")
}
