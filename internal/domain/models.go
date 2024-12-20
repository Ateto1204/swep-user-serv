package domain

import "time"

type User struct {
	ID       string    `json:"id"`
	Profile  string    `json:"profile"` // profile url
	Name     string    `json:"name"`
	Chats    []string  `json:"chats"`    // []Chat.ID
	Friends  []string  `json:"friends"`  // []User.ID
	Notifs   []string  `json:"notifs"`   // []Notification.ID
	Settings []string  `json:"settings"` // []string
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

func NewUser(ID, name string, t time.Time) *User {
	return &User{
		ID:       ID,
		Profile:  "",
		Name:     name,
		Chats:    []string{},
		Friends:  []string{},
		Notifs:   []string{},
		Settings: []string{},
		CreateAt: t,
		UpdateAt: t,
	}
}
