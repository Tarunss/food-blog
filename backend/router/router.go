package router

import (
	"net/http"

	"github.com/Tarunss/food-blog/auth"
	"github.com/Tarunss/food-blog/controller"
	"github.com/gorilla/mux"
)

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// 	// We can also check the origin of our connection with CheckOrigin, which allows us to
// 	// Make connections from our react server
// 	CheckOrigin: func(r *http.Request) bool { return true },
// }

// // Define a Reader function that will listen for connections on that particular web socket
// func reader(conn *websocket.Conn) {
// 	for {
// 		messageType, p, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		//write message out in our main
// 		fmt.Println(string(p))
// 		//Write erros if there are any from
// 		if err := conn.WriteMessage(messageType, p); err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}
// }

// Define our web socket endpoint
// func serveWs(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println(r.Host)

// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	//listen indefinitely for new messages coming in
// 	//through our web socket connection
// 	reader(ws)
// }

// Mux router (used for handling API calls)
func Router() *mux.Router {
	router := mux.NewRouter()
	//router.HandleFunc("/ws", serveWs)
	//router.HandleFunc("/api/register", controller.Register).Methods("POST")
	router.HandleFunc("/api/login", controller.PostToken).Methods("POST")
	router.HandleFunc("/api/posts", controller.GetAllPosts).Methods("GET")
	router.Handle("/api/post", auth.AuthMiddleware(http.HandlerFunc(controller.InsertOnePost))).Methods("POST")
	router.Handle("/api/post/{id}", auth.AuthMiddleware(http.HandlerFunc(controller.UpdateOnePost))).Methods("PUT")
	router.Handle("/api/post/{id}", auth.AuthMiddleware(http.HandlerFunc(controller.DeleteOnePost))).Methods("DELETE")
	router.Handle("/api/deleteallpost", auth.AuthMiddleware(http.HandlerFunc(controller.DeleteAllPosts))).Methods("DELETE")

	return router
}
