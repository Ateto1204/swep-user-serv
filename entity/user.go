package entity

import (
	"time"
)

type User struct {
	ID       string    `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name"`
	Chats    string    `json:"chats"`   // []Chat.ID
	Friends  string    `json:"friends"` // []User.ID
	Notifs   string    `json:"notifs"`  // []Notification.ID
	CreateAt time.Time `json:"create_at"`
}
