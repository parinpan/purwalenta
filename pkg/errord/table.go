package errord

import (
	"net/http"
)

var lookupTable = []ErrorComponent{
	{
		Type:       ErrFindExistingUserOnUserSignUp,
		HttpStatus: http.StatusInternalServerError,
		Message:    "",
	},
	{
		Type:       ErrUserCreationOnUserSignUp,
		HttpStatus: http.StatusInternalServerError,
		Message:    "",
	},
	{
		Type:       ErrNoAccountMatchOnUserLogin,
		HttpStatus: http.StatusInternalServerError,
		Message:    "",
	},
	{
		Type:       ErrNoMatchPasswordOnUserLogin,
		HttpStatus: http.StatusInternalServerError,
		Message:    "",
	},
}
