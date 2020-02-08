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
		Type          Error
		HttpStatus    int
		Message       string
		ServerMessage string
	}
)

const (
	ErrFindExistingUserOnUserSignUp = "ErrFindExistingUserOnUserSignUp"
	ErrUserCreationOnUserSignUp     = "ErrUserCreationOnUserSignUp"

	ErrNoAccountMatchOnUserLogin  = "ErrNoAccountMatchOnUserLogin"
	ErrNoMatchPasswordOnUserLogin = "ErrNoMatchPasswordOnUserLogin"
)
