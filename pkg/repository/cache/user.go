package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/purwalenta/purwalenta/pkg/entity"
)

const (
	signUpVerificationCode    = "verification_code_%s"
	signUpVerificationCodeTTL = 15 * time.Minute
)

type UserCacheRepository struct {
	DB *redis.Client
}

func (o *UserCacheRepository) SetSignUpVerificationCode(ctx echo.Context, verification entity.SignUpVerification) (entity.SignUpVerification, error) {
	key := fmt.Sprintf(signUpVerificationCode, verification.User.Email)
	bytes, err := json.Marshal(verification)

	if nil != err {
		return entity.SignUpVerification{}, err
	}

	if err = o.DB.Set(key, bytes, 0).Err(); nil != err {
		return entity.SignUpVerification{}, err
	}

	resp := entity.SignUpVerification{}
	resp.User.ID = verification.User.ID
	resp.User.Username = verification.User.Username
	resp.User.FullName = verification.User.FullName
	resp.User.PhoneNumber = verification.User.PhoneNumber
	resp.User.Email = verification.User.Email
	resp.User.Token = verification.User.Token
	resp.VerificationCode = verification.VerificationCode
	resp.ExpiredAt = time.Now().Add(signUpVerificationCodeTTL)

	return resp, nil
}

func (o *UserCacheRepository) GetSignUpVerificationCode(ctx echo.Context, verification entity.SignUpVerification) (entity.SignUpVerification, error) {
	key := fmt.Sprintf(signUpVerificationCode, verification.User.Email)

	var err error
	var value []byte
	var resp = entity.SignUpVerification{}

	if value, err = o.DB.Get(key).Bytes(); nil != err {
		return resp, err
	}

	if err = json.Unmarshal(value, &resp); nil != err {
		return resp, err
	}

	return resp, nil
}
