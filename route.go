package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/mohammadinasab-dev/go-websocket-chat-server/chatserver"
)

// AddApproutes will add the routes for the application
func SetupRoute(route *mux.Router) {

	log.Println("Loadeding Routes...")

	hub := chatserver.NewHub()
	go hub.Run()

	route.HandleFunc("/chat/{username}", func(responseWriter http.ResponseWriter, request *http.Request) {
		var upgrader = websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}

		// Reading username from request parameter
		username := mux.Vars(request)["username"]

		// Upgrading the HTTP connection socket connection
		connection, err := upgrader.Upgrade(responseWriter, request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		chatserver.CreateNewSocketUser(hub, connection, username)

	})

	log.Println("Routes are Loaded.")
}
