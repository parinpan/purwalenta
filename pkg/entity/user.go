package entity

import (
	"time"
)

type User struct {
	ID          string
	Username    string
	Email       string
	Password    string
	PhoneNumber string
	LoginInfo   UserLoginInfo
}

type UserLoginInfo struct {
	Token     string
	LastLogin time.Time
}

type Student struct {
	User
	Balance Balance
	Mentors []Mentor
}

type Mentor struct {
	User
	Balance  Balance
	Students []Student
}
