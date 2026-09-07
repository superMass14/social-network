package server

import (
	db "backend/database"
	"backend/app/routes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Server struct {
	Host string
	Port int
	DB   *db.Database
}

func NewServer() *Server {
	return &Server{
		Host: "0.0.0.0",
		Port: 8080,
	}
}
func (s *Server) Run() {
	fmt.Printf("Server running on port %v\n", s.Port)
	routes := routes.Route()
	file, err := os.OpenFile("server.log", os.O_APPEND|os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Configurer le logger pour écrire dans le fichier
	log.SetOutput(file)

	// Écrire des messages de débogage
	log.Printf("Hi big dev !\nlet's debug session started at : %v\n", time.Now().Format("2006/01/02 15:04:05"))
	http.ListenAndServe(s.Host+":"+strconv.Itoa(s.Port), routes)
	// Ouvrir ou créer un fichier pour écrire les logs

	fmt.Println("Server running on port", s.Port)
}
