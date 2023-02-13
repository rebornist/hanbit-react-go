package types

import "time"

type User struct {
	ID          uint      `json:"id"`
	Username    string    `json:"username"`
	UID         string    `json:"uid"`
	Password    string    `json:"password"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Age         uint8     `json:"age"`
	Birthday    string    `json:"birthday"`
	PhoneNumber string    `json:"phone_number"`
	Grade       uint8     `json:"grade"`
	Mobile      string    `json:"mobile"`
	Photo       string    `json:"photo"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
