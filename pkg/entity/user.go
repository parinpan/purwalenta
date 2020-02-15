package entity

import (
	"time"
)

const (
	DeletedUser  = 0
	ActiveUser   = 1
	InactiveUser = 2
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
	Status             int
	Type               int
}

type UserLoginInfo struct {
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
