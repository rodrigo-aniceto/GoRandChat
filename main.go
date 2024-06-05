package main

import (
	"fmt"
	"net/http"

	"GoRandChat/router"
	"GoRandChat/server"
)

func main() {

	mux := router.CreateMux()
	server.StartServer()

	fmt.Println("Server is running on :5000")
	if err := http.ListenAndServe(":5000", mux); err != nil {
		fmt.Println(err.Error())
	}
}
