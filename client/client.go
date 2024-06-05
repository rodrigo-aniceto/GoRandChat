package client

import (
	"GoRandChat/infos"
	"GoRandChat/model"
	"GoRandChat/server"
	"GoRandChat/utils"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	UserID      string
	channel     chan model.ServerMessage
	conn        *websocket.Conn
	otherUserId string
}

func (c *Client) Create(conn *websocket.Conn) {
	c.conn = conn
	c.UserID = utils.GenerateUserID(15)
	c.channel = make(chan model.ServerMessage)

	fmt.Println("UserId: ", c.UserID)
}

func (c *Client) Destroy() {
	_ = c.conn.Close()
	//close(c.channel)
}

func (c *Client) LookforChat() error {
	infos.SendInfo(c.conn, 1, "Looking for a chat...")

	server.AskConnection(c.UserID, c.channel)

	ans, open := <-c.channel

	if !open || ans.Code != 1 {
		infos.SendInfo(c.conn, 1, "Error - server connection refused!")
		return errors.New("server connection refused")
	}

	infos.SendInfo(c.conn, 2, "Chat started!")
	c.otherUserId = ans.User
	return nil

}

func (c *Client) ReceiveMessages() error {

	messageBack, open := <-c.channel

	if !open || messageBack.Code == 3 {
		infos.SendInfo(c.conn, 1, messageBack.Text)
		return errors.New("chat closed externally")
	}

	inMessage := model.UserMessage{
		Code: 3,
		User: messageBack.User,
		Text: messageBack.Text,
	}

	messageJson, err := json.Marshal(inMessage)
	if err != nil {
		fmt.Println("Error encoding message in JSON")
		return err
	}

	if err := c.conn.WriteMessage(websocket.TextMessage, messageJson); err != nil {
		return err
	}
	return nil
}

func (c *Client) SendMessages() error {

	_, messageData, err := c.conn.ReadMessage()
	if err != nil {
		infos.SendInfo(c.conn, 1, "Connection error!")
		server.EndChat(c.UserID, c.otherUserId)
		return err
	}

	var messageContent model.UserMessage
	if err = json.Unmarshal(messageData, &messageContent); err != nil {
		infos.SendInfo(c.conn, 1, "Connection error!")
		server.EndChat(c.UserID, c.otherUserId)
		return err
	}

	if err = messageContent.Prepare(); err != nil {
		infos.SendInfo(c.conn, 1, "Connection error!")
		server.EndChat(c.UserID, c.otherUserId)
		return err
	}

	if messageContent.Code == 1 {
		infos.SendInfo(c.conn, 1, "chat closed")
		server.EndChat(c.UserID, c.otherUserId)
		return errors.New("chat closed by user")
	} else if messageContent.Code == 3 {
		outMessage := model.ServerMessage{
			Code:      2,
			User:      messageContent.User,
			Text:      messageContent.Text,
			Timestamp: messageContent.Timestamp,
		}

		fmt.Println("message sent:", outMessage)
		server.SendMessage(c.otherUserId, outMessage)
	}

	return nil
}
