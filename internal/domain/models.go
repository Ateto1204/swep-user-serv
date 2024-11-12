package domain

import "time"

type User struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Chats    []string  `json:"chats"`   // []Chat.ID
	Friends  []string  `json:"friends"` // []User.ID
	CreateAt time.Time `json:"create_at"`
}
