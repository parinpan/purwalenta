package database

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/parinpan/purwalenta/pkg/entity"
	"github.com/parinpan/purwalenta/pkg/repository/database/query"
)

type UserRepository struct {
	DB *sqlx.DB
}

func (o *UserRepository) Login(ctx echo.Context, user entity.User) (*entity.User, error) {
	var result = new(entity.User)
	var row = o.DB.QueryRow(o.DB.Rebind(query.UserLoginQuery), user.Username, user.Email, user.PhoneNumber)

	var resultVars = []interface{}{
		&result.ID, &result.FullName, &result.Username, &result.Email, &result.Password, &result.PhoneNumber,
		&result.DateOfBirth, &result.Balance, &result.ProfilePicture, &result.ProfileDescription,
		&result.RefreshToken, &result.Status, &result.Type,
	}

	if err := row.Scan(resultVars...); nil != err && err != sql.ErrNoRows {
		return nil, err
	}

	return result, nil
}

func (o *UserRepository) SignUp(ctx echo.Context, user entity.User) (bool, error) {
	result := o.DB.MustExec(
		o.DB.Rebind(query.UserSignUpQuery),
		user.ID, user.FullName, user.Username, user.Email, user.Password, user.PhoneNumber,
		user.ProfilePicture, user.RefreshToken, user.Status, user.Type,
	)

	// get affected rows total and its error if exists
	numRowsAffected, err := result.RowsAffected()

	return numRowsAffected > 0, err
}

func (o *UserRepository) FindExistingUser(ctx echo.Context, user entity.User) (*entity.User, error) {
	var result = new(entity.User)

	var row = o.DB.QueryRow(
		o.DB.Rebind(query.FindExistingUserQuery),
		user.Username, user.Email, user.PhoneNumber,
	)

	var resultVars = []interface{}{
		&result.ID, &result.FullName, &result.Username, &result.Email, &result.Password, &result.PhoneNumber,
		&result.DateOfBirth, &result.Balance, &result.ProfilePicture, &result.ProfileDescription,
		&result.RefreshToken, &result.Status, &result.Type,
	}

	if err := row.Scan(resultVars...); nil != err && err != sql.ErrNoRows {
		return result, err
	}

	return result, nil
}

func (o *UserRepository) Verify(ctx echo.Context, user entity.User) (bool, error) {
	var result = o.DB.MustExec(o.DB.Rebind(query.VerifyUserSignUpQuery), entity.ActiveUser, user.Email)

	// get affected rows total and its error if exists
	numRowsAffected, err := result.RowsAffected()

	return numRowsAffected > 0, err
}

func (o *UserRepository) ChangePassword(ctx echo.Context, user entity.User) (bool, error) {
	var result = o.DB.MustExec(o.DB.Rebind(query.UserChangePasswordQuery), user.Password, user.ID)

	// get affected rows total and its error if exists
	numRowsAffected, err := result.RowsAffected()

	return numRowsAffected > 0, err
}
