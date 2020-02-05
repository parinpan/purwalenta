package database

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/purwalenta/purwalenta/pkg/entity"
	"github.com/purwalenta/purwalenta/pkg/repository/database/query"
)

type UserRepository struct {
	DB *sqlx.DB
}

func (o *UserRepository) Login(ctx echo.Context, user entity.User) (*entity.User, error) {
	return nil, nil
}

func (o *UserRepository) SignUp(ctx echo.Context, user entity.User) (bool, error) {
	result := o.DB.MustExec(
		o.DB.Rebind(query.UserSignUpQuery),
		user.ID, user.FullName, user.Username, user.Email, user.Password, user.PhoneNumber, user.Type,
	)

	numRowsAffected, err := result.RowsAffected()
	return numRowsAffected > 0, err
}

func (o *UserRepository) FindUserForSignUp(ctx echo.Context, user entity.User) (*entity.User, error) {
	var result = new(entity.User)
	var row = o.DB.QueryRow(o.DB.Rebind(query.FindUserForSignUpQuery), user.Username, user.Email, user.PhoneNumber)

	if err := row.Scan(&result.Username, &result.Email, &result.PhoneNumber); nil != err && err != sql.ErrNoRows {
		return nil, err
	}

	return result, nil
}
