package user

import "time"

type User struct {
	User_id  uint32
	Email    string
	Password string
	Role     string
}

type BasicUser struct {
	User_id uint32
	Role    string
}

type UserSession struct {
	Sid     string
	User_id uint32
	Role    string
}

type UserSessionStore struct {
	User_id uint32    `json:"userid"`
	Role    string    `json:"role"`
	Time    time.Time `json:"time"`
}
