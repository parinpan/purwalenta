package rest

import (
	"log"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	gommonLog "github.com/labstack/gommon/log"
	"github.com/parinpan/purwalenta/pkg/delivery/rest/response"
	"github.com/parinpan/purwalenta/pkg/delivery/rest/route"
	"github.com/parinpan/purwalenta/pkg/delivery/rest/validator"
	"github.com/tylerb/graceful"
)

const (
	gracefulShutdownTimeout = 5 * time.Second
)

func Start(config response.Configuration) {
	e := echo.New()
	defer e.Close()

	e = route.GetRoutes(e)
	e.Use(middleware.Recover())

	e.Validator = validator.NewRequestValidator()
	e.Server.Addr = config.Address
	e.Server.IdleTimeout = config.IdleTimeout
	e.Server.ReadTimeout = config.ReadTimeout

	file, _ := os.OpenFile("/var/log/purwalenta/error.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	defer file.Close()

	e.Logger.SetLevel(gommonLog.ERROR)
	e.Logger.SetOutput(file)

	log.Printf("Purwalenta App is now starting at %s", e.Server.Addr)
	e.Logger.Fatal(graceful.ListenAndServe(e.Server, gracefulShutdownTimeout))
}
