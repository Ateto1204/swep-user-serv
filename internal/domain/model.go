package domain

import "time"

type User struct {
	ID       string
	Name     string
	Friends  []string // User.ID
	CreateAt time.Time
}
