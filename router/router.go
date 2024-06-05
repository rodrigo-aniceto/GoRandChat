package router

import (
	"net/http"
)

func CreateMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/chat-room", chatRoomHandler)
	mux.HandleFunc("/ws/chat", chatConnectionHandler)

	fileServer := http.FileServer(http.Dir("./assets/"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fileServer))

	return mux
}
