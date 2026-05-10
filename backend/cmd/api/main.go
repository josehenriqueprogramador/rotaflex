package main

import (
	"log"

	"rotasflex/internal/handler"
)

func main() {

	server := handler.NewServer()

	log.Println("🚀 RotaFlex PRO rodando na porta 8000")

	server.Run()
}
