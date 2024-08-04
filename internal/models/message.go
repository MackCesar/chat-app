package models

// Message represents a chat message
type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}
