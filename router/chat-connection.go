package router

import (
	"GoRandChat/client"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // TODO: check origin before opening websocket
	},
}

func chatConnectionHandler(rw http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	var client client.Client

	client.Create(conn)
	//defer client.Destroy()

	if err := client.LookforChat(); err != nil {
		fmt.Println(client.UserID, err)
		return
	}

	go func() {
		for {
			if err := client.ReceiveMessages(); err != nil {
				fmt.Println(client.UserID, err)
				return
			}
		}
	}()

	for {
		if err := client.SendMessages(); err != nil {
			fmt.Println(client.UserID, err)
			break
		}
	}
	fmt.Println(client.UserID, "ending chat loop")

}
