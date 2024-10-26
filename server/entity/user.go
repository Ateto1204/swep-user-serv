package entity

import "time"

type User struct {
	ID       string    `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"create_at"`
}
