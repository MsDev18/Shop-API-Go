package entity

import "time"

type User struct {
	ID          uint
	Name        string
	Avatar      string
	PhoneNumber string
	Password    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
