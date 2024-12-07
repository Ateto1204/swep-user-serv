package domain

import "time"

type User struct {
	ID       string    `json:"id"`
	Profile  string    `json:"profile"` // profile url
	Name     string    `json:"name"`
	Chats    []string  `json:"chats"`   // []Chat.ID
	Friends  []string  `json:"friends"` // []User.ID
	Notifs   []string  `json:"notifs"`  // []Notification.ID
	CreateAt time.Time `json:"create_at"`
}

func NewUser(ID, name string, t time.Time) *User {
	return &User{
		ID:       ID,
		Profile:  "",
		Name:     name,
		Chats:    []string{},
		Friends:  []string{},
		Notifs:   []string{},
		CreateAt: t,
	}
}
