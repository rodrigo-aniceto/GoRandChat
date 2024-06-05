package server

import (
	"GoRandChat/model"
	"sync"
	"time"
)

type newUser struct {
	userID        string
	serverChannel chan model.ServerMessage
}

// FIFO to store unassigned users
var userFifo []newUser
var muFifo sync.Mutex

// Map for messages to users
var userChatMap map[string]chan model.ServerMessage
var muChat sync.Mutex

func StartServer() {
	userChatMap = make(map[string]chan model.ServerMessage)
	go conectionsLoopServer()
}

func AskConnection(userID string, userChannel chan model.ServerMessage) {
	u := newUser{
		userID:        userID,
		serverChannel: userChannel,
	}
	muFifo.Lock()
	defer muFifo.Unlock()
	userFifo = append(userFifo, u)
}

func conectionsLoopServer() {
	for {
		muFifo.Lock()
		// TODO: maybe arraylist is not the best way to do this
		for len(userFifo) >= 2 {
			userA := userFifo[0]
			userB := userFifo[1]

			userAid := userA.userID
			userAchannel := userA.serverChannel

			userBid := userB.userID
			userBchannel := userB.serverChannel

			muChat.Lock()
			userChatMap[userAid] = userAchannel
			userChatMap[userBid] = userBchannel
			muChat.Unlock()

			userAchannel <- model.ServerMessage{
				Code: 1,
				User: userBid,
			}

			userB.serverChannel <- model.ServerMessage{
				Code: 1,
				User: userAid,
			}

			userFifo = userFifo[2:]
		}
		muFifo.Unlock()
		time.Sleep(time.Second)
	}
}

// SendMessage send menssages using the channel
func SendMessage(targetUserId string, messageContent model.ServerMessage) {
	muChat.Lock()
	defer muChat.Unlock()
	ch, ok := userChatMap[targetUserId]
	if ok {
		ch <- messageContent
	}
}

// EndChat recive requestos to close a chat
func EndChat(userID, otherUserId string) {
	RemoveUser(userID)
	endMessage := model.ServerMessage{
		Code: 3,
		Text: "chat closed!",
	}
	SendMessage(otherUserId, endMessage)
	RemoveUser(otherUserId)
}

// RemoveUser remove a user from the map
func RemoveUser(UserId string) {
	muChat.Lock()
	defer muChat.Unlock()
	if userChatMap[UserId] != nil {
		if _, open := userChatMap[UserId]; open {
			close(userChatMap[UserId])
		}
	}
	delete(userChatMap, UserId)
}
