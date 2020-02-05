package api

import (
	_interface "github.com/purwalenta/purwalenta/pkg/interface"
	"github.com/purwalenta/purwalenta/pkg/repository/database"
	"github.com/purwalenta/purwalenta/pkg/repository/database/driver"
	servicePkg "github.com/purwalenta/purwalenta/pkg/service"
)

var (
	userInstances = make(map[Type]*UserAPI)
)

type UserAPI struct {
	Service _interface.UserService
}

func NewUserAPI(apiType Type) *UserAPI {
	if _, exists := userInstances[apiType]; exists {
		return userInstances[apiType]
	}

	switch apiType {
	case DefaultUserAPIFlag:
		return newDefaultUserAPI()
	}

	return newDefaultUserAPI()
}

func newDefaultUserAPI() *UserAPI {
	once.Do(func() {
		repo := new(database.UserRepository)
		repo.DB, _ = driver.GetPostgreDriver()

		service := new(servicePkg.UserService)
		service.Repo = repo

		userInstances[DefaultUserAPIFlag] = &UserAPI{
			Service: service,
		}
	})

	return userInstances[DefaultUserAPIFlag]
}
