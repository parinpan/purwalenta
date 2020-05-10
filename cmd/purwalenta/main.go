package main

import (
	"time"

	"github.com/parinpan/purwalenta/pkg/delivery/rest"
	"github.com/parinpan/purwalenta/pkg/delivery/rest/response"
)

func main() {
	rest.Start(response.Configuration{
		Address:     ":9099",
		IdleTimeout: 2 * time.Second,
		ReadTimeout: 2 * time.Second,
	})
}
