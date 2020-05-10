package api

import (
	_interface "github.com/parinpan/purwalenta/pkg/interface"
	"github.com/parinpan/purwalenta/pkg/repository/cache"
	cacheDriver "github.com/parinpan/purwalenta/pkg/repository/cache/driver"
	"github.com/parinpan/purwalenta/pkg/repository/database"
	"github.com/parinpan/purwalenta/pkg/repository/database/driver"
	"github.com/parinpan/purwalenta/pkg/repository/mail"
	mailingDriver "github.com/parinpan/purwalenta/pkg/repository/mail/driver"
	servicePkg "github.com/parinpan/purwalenta/pkg/service"
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

		cacheRepo := new(cache.UserCacheRepository)
		cacheRepo.DB, _ = cacheDriver.GetRedisDriver()

		mailingRepo := new(mail.UserMailingRepository)
		mailingRepo.Driver = mailingDriver.GoMailDriver

		service := new(servicePkg.UserService)
		service.Repo = repo
		service.CacheRepo = cacheRepo
		service.MailingRepo = mailingRepo

		userInstances[DefaultUserAPIFlag] = &UserAPI{
			Service: service,
		}
	})

	return userInstances[DefaultUserAPIFlag]
}
