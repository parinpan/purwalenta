package rest

import (
	"log"
	"time"
	
	"github.com/labstack/echo"
	"github.com/tylerb/graceful"
)

const (
	gracefulShutdownTimeout = 5 * time.Second
)

func Start(config Configuration) {
	e := echo.New()
	e = registerRoutes(e)
	e.Server.Addr = config.Address
	
	log.Printf("Purwalenta App is now starting at %s", e.Server.Addr)
	e.Logger.Fatal(graceful.ListenAndServe(e.Server, gracefulShutdownTimeout))
}

