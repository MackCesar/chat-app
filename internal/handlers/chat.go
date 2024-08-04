package handlers

import (
	"chat-app/internal/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// SetupRouter initialize the router and routes
func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/chat", HandleChat).Methods("POST")
	return router
}

func HandleChat(w http.ResponseWriter, r *http.Request) {
	var msg models.Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	response := models.Message{
		Username: msg.Username,
		Content:  "You said: " + msg.Content,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
