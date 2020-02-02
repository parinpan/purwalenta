package main

import (
	"github.com/purwalenta/purwalenta/pkg/delivery/rest"
)

func main() {
	rest.Start(rest.Configuration{
		Address: ":9099",
	})
}
