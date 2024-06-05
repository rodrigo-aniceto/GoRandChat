package infos

import (
	"GoRandChat/model"
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

// SendInfo sends a message to the front end, code 1 for blocking messages, code 2 for non-blocking messages
func SendInfo(conn *websocket.Conn, code int, msg string) {
	messageBack := model.UserMessage{
		Code: code,
		Text: msg,
	}
	messageBack.Prepare()

	messageJson, err := json.Marshal(messageBack)
	if err != nil {
		fmt.Println("Error encoding message in JSON:", err)
		return
	}
	if err := conn.WriteMessage(websocket.TextMessage, messageJson); err != nil {
		fmt.Println(err)
		return
	}
}
