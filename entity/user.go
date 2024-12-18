package entity

import (
	"time"
)

type User struct {
	ID       string    `gorm:"primaryKey" json:"id"`
	Profile  string    `json:"profile"` // profile url
	Name     string    `json:"name"`
	Chats    string    `json:"chats"`    // []Chat.ID
	Friends  string    `json:"friends"`  // []User.ID
	Notifs   string    `json:"notifs"`   // []Notification.ID
	Settings string    `json:"settings"` // []string
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `gorm:"autoUpdateTime" json:"update_at"`
}
