package entities

import "time"

type User struct {
	ID          uint
	Firstname   string
	Lastname    string
	Email       string
	PhoneNumber uint
	Role        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
