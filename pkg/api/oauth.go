package api

import (
	_interface "github.com/parinpan/purwalenta/pkg/interface"
	"github.com/parinpan/purwalenta/pkg/repository/apicalls"
	"github.com/parinpan/purwalenta/pkg/repository/database"
	"github.com/parinpan/purwalenta/pkg/repository/database/driver"
	servicePkg "github.com/parinpan/purwalenta/pkg/service"
)

var (
	oauthInstances = make(map[Type]*OauthAPI)
)

type OauthAPI struct {
	Service _interface.OauthService
}

func NewOauthAPI(apiType Type) *OauthAPI {
	if _, exists := oauthInstances[apiType]; exists {
		return oauthInstances[apiType]
	}

	switch apiType {
	case DefaultOauthAPIFlag:
		return newDefaultOauthAPI()
	}

	return newDefaultOauthAPI()
}

func newDefaultOauthAPI() *OauthAPI {
	userRepo := new(database.UserRepository)
	userRepo.DB, _ = driver.GetPostgreDriver()

	service := new(servicePkg.OauthService)
	service.UserRepo = userRepo
	service.GoogleOauthRepo = new(apicalls.GoogleOauthRepository)

	oauthInstances[DefaultOauthAPIFlag] = &OauthAPI{
		Service: service,
	}

	return oauthInstances[DefaultOauthAPIFlag]
}
