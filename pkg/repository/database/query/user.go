package query

const (
	FindExistingUserQuery = `
		SELECT
			username, email, phone_number
		FROM
			public.user
		WHERE
			username = ? OR email = ? OR phone_number = ?
	`

	UserLoginQuery = `
		SELECT
			*
		FROM
			public.user
		WHERE
			username = ? OR email = ? OR phone_number = ?
	`

	UserSignUpQuery = `
		INSERT INTO
			public.user(id, full_name, username, email, password, phone_number, profile_picture, refresh_token, status, type)
		VALUES (
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?
		)
	`
)
