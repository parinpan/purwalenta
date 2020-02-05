package query

const (
	FindUserForSignUpQuery = `
		SELECT
			username, email, phone_number
		FROM
			public.user
		WHERE
			username = ? OR email = ? OR phone_number = ?
	`

	UserSignUpQuery = `
		INSERT INTO
			public.user(id, full_name, username, email, password, phone_number, type)
		VALUES (
			?, ?, ?, ?, ?, ?, ?
		)
	`
)
