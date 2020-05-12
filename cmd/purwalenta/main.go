package main

import (
	"time"
	
	"github.com/parinpan/purwalenta/pkg/delivery/rest"
)

func main() {
	rest.Start(rest.Configuration{
		Address:     ":9099",
		IdleTimeout: 2 * time.Second,
		ReadTimeout: 2 * time.Second,
	})
}
