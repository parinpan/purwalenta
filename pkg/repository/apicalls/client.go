package apicalls

import (
	"sync"

	"github.com/go-resty/resty"
)

var (
	once                   sync.Once
	apicallsClientInstance *resty.Client
)

func getClient() *resty.Client {
	once.Do(func() {
		apicallsClientInstance = resty.New()
	})

	return apicallsClientInstance
}
