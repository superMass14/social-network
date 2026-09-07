package main

import (
	server "backend/server"
)

func main() {
	server := server.NewServer()
	server.Run()
}
