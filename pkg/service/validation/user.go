package validation

import (
	"github.com/purwalenta/purwalenta/pkg/entity"
)

func ValidateUserSignUpTakenFields(user *entity.User) (takenFields []string, isTaken bool) {
	if nil == user {
		return takenFields, isTaken
	}

	var fieldsToValidate = map[string]string{
		"username":     user.Username,
		"email":        user.Email,
		"phone_number": user.PhoneNumber,
	}

	for field, value := range fieldsToValidate {
		if value != "" {
			isTaken = true
			takenFields = append(takenFields, field)
		}
	}

	return takenFields, isTaken
}
