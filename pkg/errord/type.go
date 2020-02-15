package errord

import (
	"sync"
)

var (
	once              sync.Once
	lookupMapInstance = make(map[Error]ErrorComponent)
)

type (
	Error string

	ErrorComponent struct {
		Type          Error  `json:"type"`
		HttpStatus    int    `json:"http_status"`
		Message       string `json:"message"`
		ServerMessage string `json:"server_message"`
	}

	Option struct {
		WriteLog       bool
		FormatterValue []interface{}
	}

	errorOutput struct {
		System     string         `json:"system"`
		Dictionary ErrorComponent `json:"errord"`
	}
)

const (
	ErrFindExistingUserOnUserSignUp = "ErrFindExistingUserOnUserSignUp"
	ErrUserCreationOnUserSignUp     = "ErrUserCreationOnUserSignUp"
	ErrFieldHasTakenOnUserSignUp    = "ErrFieldHasTakenOnUserSignUp"

	ErrGeneralOnUserVerify     = "ErrInvalidCodeOnUserVerify"
	ErrInvalidCodeOnUserVerify = "ErrInvalidCodeOnUserVerify"

	ErrGeneralOnUserSendVerificationCode = "ErrGeneralOnUserSendVerificationCode"

	ErrNoAccountMatchOnUserLogin  = "ErrNoAccountMatchOnUserLogin"
	ErrNoMatchPasswordOnUserLogin = "ErrNoMatchPasswordOnUserLogin"
)
