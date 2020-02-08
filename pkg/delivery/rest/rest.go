package rest

import (
	"log"
	"time"

	"github.com/labstack/echo"
	"github.com/purwalenta/purwalenta/pkg/delivery/rest/response"
	"github.com/purwalenta/purwalenta/pkg/delivery/rest/route"
	"github.com/purwalenta/purwalenta/pkg/delivery/rest/validator"
	"github.com/tylerb/graceful"
)

const (
	gracefulShutdownTimeout = 5 * time.Second
)

func Start(config response.Configuration) {
	e := echo.New()
	defer e.Close()

	e = route.GetRoutes(e)
	e.Validator = validator.NewRequestValidator()
	e.Server.Addr = config.Address
	e.Server.IdleTimeout = config.IdleTimeout
	e.Server.ReadTimeout = config.ReadTimeout

	log.Printf("Purwalenta App is now starting at %s", e.Server.Addr)
	e.Logger.Fatal(graceful.ListenAndServe(e.Server, gracefulShutdownTimeout))
}
