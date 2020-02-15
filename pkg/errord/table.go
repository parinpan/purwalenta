package errord

import (
	"net/http"
)

var lookupTable = []ErrorComponent{
	{
		Type:       ErrFindExistingUserOnUserSignUp,
		HttpStatus: http.StatusInternalServerError,
		Message:    "Duh... sedang terjadi kesalahan. Coba kembali beberapa saat lagi ya.",
	},
	{
		Type:       ErrUserCreationOnUserSignUp,
		HttpStatus: http.StatusInternalServerError,
		Message:    "Duh... sedang terjadi kesalahan. Coba kembali beberapa saat lagi ya.",
	},
	{
		Type:       ErrFieldHasTakenOnUserSignUp,
		HttpStatus: http.StatusBadRequest,
		Message:    "Yah... identitas %s kamu sudah pernah terdaftar. Coba gunakan yang lain ya.",
	},
	{
		Type:       ErrInvalidCodeOnUserVerify,
		HttpStatus: http.StatusBadRequest,
		Message:    "Kode verifikasi yang kamu masukkan salah. Coba lagi ya.",
	},
	{
		Type:       ErrGeneralOnUserSendVerificationCode,
		HttpStatus: http.StatusInternalServerError,
		Message:    "Duh... sedang terjadi kesalahan. Coba kembali beberapa saat lagi ya.",
	},
	{
		Type:       ErrNoAccountMatchOnUserLogin,
		HttpStatus: http.StatusBadRequest,
		Message:    "Pastikan username dan password kamu sudah benar ya.",
	},
	{
		Type:       ErrNoMatchPasswordOnUserLogin,
		HttpStatus: http.StatusBadRequest,
		Message:    "Pastikan username dan password kamu sudah benar ya.",
	},
}
