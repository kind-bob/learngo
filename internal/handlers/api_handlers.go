package handlers

import (
	"encoding/json"
	"learngo/internal/messagesService" // Импортируем наш сервис
	"net/http"
)

type Handler struct {
	Service *messagesService.MessageService
}

// Нужна для создания структуры Handler на этапе инициализации приложения

func NewHandler(service *messagesService.MessageService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	messages, err := h.Service.GetAllMessages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func (h *Handler) PostMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message messagesService.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdMessage, err := h.Service.CreateMessage(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdMessage)
}

func (h *Handler) PatchMessagesHandler(w http.ResponseWriter, r *http.Request) {
	var newString messagesService.Message

	err := json.NewDecoder(r.Body).Decode(&newString)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	UpdatedMessage, err := h.Service.UpdateMessageById(newString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(UpdatedMessage)
}

func (h *Handler) DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	var deleteString messagesService.Message
	err := json.NewDecoder(r.Body).Decode(&deleteString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	errr := h.Service.DeleteMessageById(deleteString)
	if errr != nil {
		http.Error(w, errr.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(errr)
}
