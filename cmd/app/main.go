package main

import (
	"learngo/internal/database"
	"learngo/internal/handlers"
	"learngo/internal/messagesService"
	"net/http"

	"github.com/gorilla/mux"
)

// Переменная для работы с БД

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&messagesService.Message{})

	repo := messagesService.NewMessageRepository(database.DB)
	service := messagesService.NewService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetMessagesHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostMessageHandler).Methods("POST")
	router.HandleFunc("/api/update", handler.PatchMessagesHandler).Methods("PATCH")
	http.ListenAndServe(":8080", router)
}
