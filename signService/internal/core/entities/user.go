package entities

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID          uint
	Firstname   string
	Lastname    string
	Email       string
	PhoneNumber uint
	Password    string
	Role        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (user *User) HashPassword(password string) error {
	fmt.Println(password)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
