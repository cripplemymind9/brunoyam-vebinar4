package main

import (
	"github.com/cripplemymind9/brunoyam-vebinar4/homework/internal/storage"
	"github.com/cripplemymind9/brunoyam-vebinar4/homework/internal/server"
	"log"
)

func main() {
	store, err := storage.New()
	if err != nil {
		log.Fatal(err)
	}

	server := server.NewServer(":8080", store)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}