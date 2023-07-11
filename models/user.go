package models

import "time"

type User struct {
	ID        int64     `json:"id"`
	Fio       string    `json:"fio"`
	Number    string    `json:"number"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
}
