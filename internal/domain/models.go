package domain

import "time"

type User struct {
	ID       string
	Name     string
	Chats    []string // []Chat.ID
	Friends  []string // []User.ID
	CreateAt time.Time
}
