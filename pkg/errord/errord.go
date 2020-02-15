package errord

import (
	"errors"
	"fmt"

	"github.com/labstack/echo"
)

func New(ctx echo.Context, systemError error) func(Error, ...Option) error {
	ctx.Set("system_error", systemError)
	var systemErrorMsg string

	if nil != systemError {
		systemErrorMsg = systemError.Error()
	}

	return func(errordKey Error, options ...Option) error {
		option := getOption(options...)
		errordError := getErrorOnLookup(errordKey)

		if errordError.Type != "" {
			errordError.Message = fmt.Sprintf(errordError.Message, option.FormatterValue...)
			errordError.ServerMessage = systemErrorMsg
		}

		ctx.Set("errord_error", errordError)
		errorStr := fmt.Sprintf("System: %s | Errord: %+v", systemErrorMsg, errordError)

		if option.WriteLog {
			defer ctx.Logger().Error(errorStr)
		}

		return errors.New(errorStr)
	}
}

func getErrorOnLookup(key Error) ErrorComponent {
	once.Do(func() {
		for _, errComponent := range lookupTable {
			lookupMapInstance[errComponent.Type] = errComponent
		}
	})

	if errComponent, ok := lookupMapInstance[key]; ok {
		return errComponent
	}

	return ErrorComponent{}
}

func getOption(options ...Option) Option {
	var opt Option

	for _, option := range options {
		opt.WriteLog = option.WriteLog
		opt.FormatterValue = option.FormatterValue
	}

	return opt
}
