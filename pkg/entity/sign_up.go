package entity

import (
	"time"
)

type SignUpVerification struct {
	User
	VerificationCode int
	Token            string
	ExpiredAt        time.Time
}
