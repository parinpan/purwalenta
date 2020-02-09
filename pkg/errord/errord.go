package errord

import (
	"github.com/labstack/echo"
)

func New(ctx echo.Context, err error) func(Error) error {
	ctx.Set("server_error", err)

	return func(errordKey Error) error {
		ctx.Set("errord_error", getErrorOnLookup(errordKey))
		return nil
	}
}

func getErrorOnLookup(key Error) ErrorComponent {
	once.Do(func() {
		for _, errComponent := range lookupMapInstance {
			lookupMapInstance[errComponent.Type] = errComponent
		}
	})

	if errComponent, ok := lookupMapInstance[key]; ok {
		return errComponent
	}

	return ErrorComponent{}
}
