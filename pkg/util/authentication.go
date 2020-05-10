package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/parinpan/purwalenta/pkg/config"
	"github.com/parinpan/purwalenta/pkg/entity"
)

type jwtClaims struct {
	jwt.StandardClaims
	ID          string `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func GenerateUserLoginToken(cfg config.Config, user entity.User) (string, error) {
	var timeNow = time.Now()
	var expiredAt = timeNow.Add(time.Microsecond * time.Duration(cfg.UserAuthentication.MaxLifeTime))

	claims := jwtClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    cfg.App.Name,
			IssuedAt:  timeNow.Unix(),
			ExpiresAt: expiredAt.Unix(),
		},
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(cfg.UserAuthentication.SecretToken))

	if nil != err {
		return "", err
	}

	return signedToken, nil
}
