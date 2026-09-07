package ws

import (
	// "backend/server/handler"

	"backend/app/handler/groups/events"
	joingroup "backend/app/handler/groups/joinGroup"
	"backend/app/service"
	"backend/database"
	"backend/utils/seed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	WS_JOIN_EVENT             = "join-event"
	WS_DISCONNECT_EVENT       = "disconnect-event"
	WS_IDRECEIVER_EVENT       = "id-receiver-event"
	WS_MESSAGEUSER_EVENT      = "message-user-event"
	WS_MESSAGEGROUP_EVENT     = "message-group-event"
	WS_IDGROUP_RECEIVER_EVENT = "idGroup-receiver-event"
	WS_NAVBAR_MESSAGE         = "message-navbar"
)

type WSClient struct {
	Email        string
	WSCoon       *websocket.Conn
	OutgoingMsg  chan interface{}
	SessionToken string
}
type WSPaylaod struct {
	From string
	Type string
	Data interface{}
	To   string
}
type Hub struct {
	Clients           *sync.Map
	RegisterChannel   chan *WSClient
	UnRegisterChannel chan *WSClient
	SSE               chan WSPaylaod
}

var WSHub *Hub

func init() {
	WSHub = newHub()
	go WSHub.listen()
}
func newHub() *Hub {
	return &Hub{
		Clients:           &sync.Map{},
		RegisterChannel:   make(chan *WSClient),
		UnRegisterChannel: make(chan *WSClient),
		SSE:               make(chan WSPaylaod),
	}
}
func (h *Hub) listen() {
	for {
		select {
		case client := <-h.RegisterChannel:
			h.Clients.Store(client.Email, client)
			log.Printf("Client %s connected\n", client.Email)
		case client := <-h.UnRegisterChannel:
			// if _, ok := h.Clients.Load(client.Email); ok {
			h.Clients.Delete(client.Email)
			close(client.OutgoingMsg)
			log.Printf("Client %s disconnected\n", client.Email)
			// }
		case message := <-h.SSE:
			h.HandleEvent(message)
		}
	}
}

var groupid int

func (h *Hub) HandleEvent(eventPayload WSPaylaod) {
	switch eventPayload.Type {
	case WS_JOIN_EVENT:
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Email != eventPayload.From {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case WS_DISCONNECT_EVENT:
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Email != eventPayload.From {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case WS_NAVBAR_MESSAGE:
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)

			if client.Email == eventPayload.From {
				client.OutgoingMsg <- eventPayload
			}

			return true
		})
	case WS_IDRECEIVER_EVENT:
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Email == eventPayload.From {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case WS_IDGROUP_RECEIVER_EVENT:
		h.Clients.Range(func(key, value any) bool {
			client := value.(*WSClient)
			if client.Email == eventPayload.From {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case WS_MESSAGEUSER_EVENT:
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Email == eventPayload.From || client.Email == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case WS_MESSAGEGROUP_EVENT:
		clientFollowers, err := seed.GetFollowerGroup(seed.DB, groupid)
		if err != nil {
			fmt.Println("error", err)
		}
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			// Vérification si le client suit le groupe
			for _, follower := range clientFollowers {
				if client.Email == follower.Email {
					client.OutgoingMsg <- eventPayload
					break
				}
			}
			return true
		})

		//===============================================

	case "Welcome":
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			// if client.Email == eventPayload.To {
			client.OutgoingMsg <- eventPayload
			// }
			return true
		})
	case "JoinGroup":
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Email == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case "NotGoingEvent":
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Email == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case "GoingEvent":
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Email == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case "Desplayed Events":
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Email == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case "SuggestFriend":
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Email == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case "follow":
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Email == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case "notification":
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			// if client.Mail == eventPayload.To {
			client.OutgoingMsg <- eventPayload
			// }
			return true
		})
	case "ResponceNotification":
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			// if client.Mail != eventPayload.To {
			client.OutgoingMsg <- eventPayload
			// }
			return true
		})
	}

}

func (wsHub *Hub) AddClient(coon *websocket.Conn, Email string, sessionToken string, r *http.Request) {
	client := &WSClient{
		Email:        Email,
		WSCoon:       coon,
		OutgoingMsg:  make(chan interface{}),
		SessionToken: sessionToken,
	}
	go client.messageReader(r)
	go client.messageWriter()
	wsHub.RegisterChannel <- client
	var newEvent = WSPaylaod{
		From: client.Email,
		Type: WS_JOIN_EVENT,
		Data: "New client joined",
	}
	wsHub.HandleEvent(newEvent)
}
func (client *WSClient) messageReader(r *http.Request) {

	userIdConnected, err := service.AuthServ.VerifyToken(r)
	if err != nil {
		fmt.Println("verify token error", err)
	}

	date := time.Now().Format("2006-01-02T15:04:05")
	Db := database.NewDatabase()
	defer Db.Close()
	for {
		_, message, err := client.WSCoon.ReadMessage()
		if err != nil {
			WSHub.UnRegisterChannel <- client
			var newEvent = WSPaylaod{
				From: client.Email,
				Type: WS_DISCONNECT_EVENT,
				Data: nil,
			}
			WSHub.HandleEvent(newEvent)
			return
		}
		var payload map[string]interface{}
		err = json.Unmarshal(message, &payload)
		if err != nil {
			return
		}
		eventType := payload["type"].(string)

		wsEvent := WSPaylaod{
			From: client.Email,
			Type: eventType,
			Data: payload,
		}

		// EvenType le type d'evenement socket
		switch eventType {

		case WS_NAVBAR_MESSAGE:
			listGroup, _ := seed.GetGroup(seed.DB, int(payload["data"].(float64)))
			listUser, err := seed.ListeUsers(seed.DB, int(payload["data"].(float64)))
			if err != nil {
				fmt.Println("error:", err)
			} else if len(listUser) == 0 {
				fmt.Println("La liste des utilisateurs est vide")
			} else {

				data := []interface{}{listUser[0], listGroup}

				var newEvent = WSPaylaod{
					From: client.Email,
					Type: eventType,
					Data: data, // Assigner le tableau à Data
				}
				WSHub.HandleEvent(newEvent)
			}

		case WS_IDRECEIVER_EVENT:
			clickedUserId, ok := payload["data"].(map[string]interface{})["clickedUserId"].(float64)
			if !ok {
				fmt.Println("Erreur lors de l'accès à la clé 'clickedUserId'")
				return
			}
			sessionUserId, ok := payload["data"].(map[string]interface{})["sessionUserId"].(float64)
			if !ok {
				fmt.Println("Erreur lors de l'accès à la clé 'sendId'")
				return
			}

			data, _ := seed.SelectMsgBetweenUsers(seed.DB, int(sessionUserId), int(clickedUserId))
			wsEvent := WSPaylaod{
				From: client.Email,
				Type: eventType,
				Data: data,
			}
			WSHub.HandleEvent(wsEvent)

		case WS_IDGROUP_RECEIVER_EVENT:

			groupId, ok := payload["data"].(map[string]interface{})["idgroup"].(float64)
			if !ok {
				fmt.Println("Erreur lors de l'accès à la clé 'clickedUserId'")
				return
			}

			data, _ := seed.GetGroupMessage(seed.DB, int(groupId))
			wsEvent := WSPaylaod{
				From: client.Email,
				Type: eventType,
				Data: data,
			}
			WSHub.HandleEvent(wsEvent)

		case WS_MESSAGEUSER_EVENT:

			message, ok := payload["data"].(map[string]interface{})["message"].(string)
			if !ok {
				fmt.Println("Erreur lors de l'accès à la clé 'message'")
				return
			}
			sendId, ok := payload["data"].(map[string]interface{})["sendId"].(float64)
			if !ok {
				fmt.Println("Erreur lors de l'accès à la clé 'sendId'")
				return
			}
			receiverID, ok := payload["data"].(map[string]interface{})["receiverId"].(float64)
			if !ok {
				fmt.Println("Erreur lors de l'accès à la clé 'sendId'")
				return
			}

			err := seed.InsertMessage(seed.DB, int(sendId), int(receiverID), message, date)
			if err != nil {
				fmt.Println(err)
			}
			mes, errr := seed.GetLastMessage(seed.DB, message)
			if errr != nil {
				fmt.Println(errr)
			}
			clientTo, _ := seed.GetUserById(seed.DB, int(receiverID))
			msEvent := WSPaylaod{
				From: client.Email,
				Type: eventType,
				Data: mes,
				To:   clientTo.Email,
			}
			WSHub.HandleEvent(msEvent)

		case WS_MESSAGEGROUP_EVENT:

			message, ok := payload["data"].(map[string]interface{})["message"].(string)
			if !ok {
				fmt.Println("Erreur lors de l'accès à la clé 'message'")
				return
			}
			sendId, ok := payload["data"].(map[string]interface{})["sendId"].(float64)
			if !ok {
				fmt.Println("Erreur lors de l'accès à la clé 'sendId'")
				return
			}
			groupReceiverID, ok := payload["data"].(map[string]interface{})["receiverId"].(float64)
			if !ok {
				fmt.Println("Erreur lors de l'accès à la clé 'sendId'")
				return
			}
			groupid = int(groupReceiverID)
			err := seed.InsertGroupMessage(seed.DB, int(sendId), int(groupReceiverID), message, date)
			if err != nil {
				fmt.Println("error", err)
			}
			lastmsg, errr := seed.GetLastGroupMessage(seed.DB, message)
			if errr != nil {
				fmt.Println("error", err)

			}
			GPEvent := WSPaylaod{
				From: client.Email,
				Type: eventType,
				Data: lastmsg,
			}
			WSHub.HandleEvent(GPEvent)

			//==================================

		case "JoinGroup":
			jsonData, err := json.Marshal(wsEvent.Data)
			if err != nil {
				fmt.Println("Erreur de conversion en json", err)
				return
			}
			var parseData map[string]interface{}
			if err := json.Unmarshal(jsonData, &parseData); err != nil {
				fmt.Println("Erreur de conversion en json", err)
			}

			groupeId, ok := parseData["groupId"].(float64)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee")
				return
			}

			typeNotification, ok := parseData["type"].(string)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee")
				return
			}
			err = joingroup.InsertNotification(int(groupeId), typeNotification, userIdConnected.User_id, 0, Db)
			if err != nil {
				fmt.Println("Error inserting", err.Error())
			}

			idAdminGroup, err := joingroup.RecupeIdAdminGroup(int(groupeId), Db)
			if err != nil {
				fmt.Println("Erreur lors de la recuperation de l'id de l'administrateur du groupe ", err)
			}
			dataSend := struct {
				IdGroup int    `json:"id_group"`
				Button  string `json:"button"`
			}{
				IdGroup: int(groupeId),
				Button:  "Desable",
			}

			wsEvent = WSPaylaod{
				From: client.Email,
				Type: eventType,
				Data: dataSend,
				To:   string(rune(idAdminGroup)),
			}
			WSHub.HandleEvent(wsEvent)

		case "NotGoingEvent":
			jsonData, err := json.Marshal(wsEvent.Data)
			if err != nil {
				fmt.Println("Erreur de conversion en json", err)
				return
			}
			var parseData map[string]interface{}
			if err := json.Unmarshal(jsonData, &parseData); err != nil {
				fmt.Println("Erreur de conversion en json", err)
			}

			groupeId, ok := parseData["groupId"].(float64)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee")
				return
			}
			event_id, ok := parseData["event_id"].(float64)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee")
				return
			}

			err = events.NotJoinEvent(userIdConnected.User_id, int(groupeId), int(event_id), Db)
			if err != nil {
				log.Fatal(err.Error())
			}

		case "GoingEvent":
			jsonData, err := json.Marshal(wsEvent.Data)
			if err != nil {
				fmt.Println("Erreur de conversion en json", err)
				return
			}
			var parseData map[string]interface{}
			if err := json.Unmarshal(jsonData, &parseData); err != nil {
				fmt.Println("Erreur de conversion en json", err)
			}

			groupeId, ok := parseData["groupId"].(float64)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee")
				return
			}
			event_id, ok := parseData["event_id"].(float64)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee")
				return
			}
			err = events.JoinEvent(userIdConnected.User_id, int(groupeId), int(event_id), Db)
			if err != nil {
				log.Fatal(err.Error())
			}
		case "SuggestFriend":
			jsonData, err := json.Marshal(wsEvent.Data)
			if err != nil {
				fmt.Println("Erreur de conversion en json", err)
				return
			}
			fmt.Println("Suggest = ", string(jsonData))
			var parseData map[string]interface{}
			if err := json.Unmarshal(jsonData, &parseData); err != nil {
				fmt.Println("Erreur de conversion en json", err)
			}

			userid, ok := parseData["userId"].(float64)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee userid")
				return
			}

			groupeID, ok := parseData["idGroup"].(string)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee groupId")
				return
			}
			typeNotification, ok := parseData["type"].(string)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee type ")
				return
			}
			group_id, err := strconv.Atoi(groupeID)
			if err != nil {
				log.Fatal(err.Error())

			}
			err = joingroup.InsertNotification(group_id, typeNotification, userIdConnected.User_id, int(userid), Db)
			if err != nil {
				fmt.Println("Error inserting notification ", err.Error())
				return
			}
		case "ResponceNotification":
			jsonData, err := json.Marshal(wsEvent.Data)
			if err != nil {
				fmt.Println("Erreur de conversion en json", err)
				return
			}
			var parseData map[string]interface{}
			if err := json.Unmarshal(jsonData, &parseData); err != nil {
				fmt.Println("Erreur de conversion en json", err)
			}

			id_user_sender, ok := parseData["id_user_sender"].(float64)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee")
				return
			}
			id_user_receiver, ok := parseData["id_user_receiver"].(float64)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee id_user_receiver !!!")
				return
			}

			groupid, ok := parseData["groupeId"].(float64)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee id_group")
				return
			}

			typeNotification, ok := parseData["typeNotification"].(string)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee typeNotification")
				return
			}

			response, ok := parseData["response"].(string)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee response")
				return
			}

			fmt.Println("Reponse ", response)
			if response == "going" {
				if typeNotification == "follow" {
					err := service.AuthServ.FollowUser(int(id_user_sender), int(id_user_receiver))
					if err != nil {
						fmt.Println("Erreur lors de l'insertion de users follower ", err)
						return
					}
					check := joingroup.AcceptOrNoSugg(Db, int(id_user_sender), int(id_user_receiver), "1")
					if !check {
						fmt.Println("Error lors de la requete update ", check)
						return
					}
				} else {

					check := joingroup.AcceptOrNo(Db, int(id_user_sender), int(id_user_receiver), int(groupid), "1")
					if !check {
						fmt.Println("Error lors de la requete update ", check)
						return
					} else {
						if typeNotification == "SuggestFriend" {
							tpm := id_user_sender
							id_user_sender = id_user_receiver
							id_user_receiver = tpm
						}
						joingroup.InsertGroupFollowers(Db, int(id_user_sender), int(groupid))
					}
				}
				joingroup.DeleteNotification(Db, "1")
			} else if response == "notGoing" {
				if typeNotification == "follow" {
					check := joingroup.AcceptOrNoSugg(Db, int(id_user_sender), int(id_user_receiver), "-1")
					if !check {
						fmt.Println("Error lors de la requete update ", check)
						return
					}
				} else {
					check := joingroup.AcceptOrNo(Db, int(id_user_sender), int(id_user_receiver), int(groupid), "-1")
					if !check {
						fmt.Println("Error lors de la requete update ", check)
						return
					}
				}
				joingroup.DeleteNotification(Db, "-1")
			}
		case "follow":
			jsonData, err := json.Marshal(wsEvent.Data)
			if err != nil {
				fmt.Println("Erreur de conversion en json", err)
				return
			}
			var parseData map[string]interface{}
			if err := json.Unmarshal(jsonData, &parseData); err != nil {
				fmt.Println("Erreur de conversion en json", err)
			}
			followerId, ok := parseData["follower_id"].(float64)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee")
				return
			}
			followedId, ok := parseData["followed_id"].(float64)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee")
				return
			}
			err = joingroup.InsertFollowNotification(int(followerId), int(followedId), Db)
			if err != nil {
				fmt.Println("Error inserting", err.Error())
			}
		}

	}
}
func (client *WSClient) messageWriter() {
	for {
		select {
		case message := <-client.OutgoingMsg:
			data, err := json.Marshal(message)
			if err != nil {
				return
			}
			err = client.WSCoon.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				return
			}
		}
	}
}
