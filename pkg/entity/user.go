package entity

import (
	"time"
)

type User struct {
	ID                 string
	FullName           string
	Username           string
	Email              string
	Password           string
	PhoneNumber        string
	DateOfBirth        *time.Time
	Balance            float64
	ProfilePicture     string
	ProfileDescription string
	LoginInfo          UserLoginInfo
	Token              string
	RefreshToken       string
	Type               int
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
