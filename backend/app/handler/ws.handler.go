package handler

import (
	"backend/app/cors"
	"backend/app/ws"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024, // Autoriser toutes les origines pour le développement
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WSHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w) // Configurez les en-têtes CORS pour les réponses HTTP
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	userSession := UserInSession(w, r)
	token := r.URL.Query().Get("token")
	conn, err := upgrader.Upgrade(w, r, nil) // Utilisez l'Upgrader configuré
	if err != nil {
		log.Println(err)
		return
	}
	// defer conn.Close()

	ws.WSHub.AddClient(conn, userSession.Email, token, r)

}
