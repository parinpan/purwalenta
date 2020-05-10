package validation

import (
	"github.com/parinpan/purwalenta/pkg/entity"
	"github.com/parinpan/purwalenta/pkg/service/request"
)

func ValidateUserSignUpTakenFields(req request.UserSignUp, user entity.User) (takenFields []string, isTaken bool) {
	var valuesToValidate = map[string]string{
		"username":     user.Username,
		"email":        user.Email,
		"phone_number": user.PhoneNumber,
	}

	var requestFieldValues = map[string]string{
		"username":     req.Username,
		"email":        req.Email,
		"phone_number": req.PhoneNumber,
	}

	for field, value := range valuesToValidate {
		if value == requestFieldValues[field] {
			isTaken = true
			takenFields = append(takenFields, field)
		}
	}

	return takenFields, isTaken
}
