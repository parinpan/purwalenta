package service

import (
	"context"
	"errors"
	"time"

	"github.com/labstack/echo"
	"github.com/parinpan/purwalenta/pkg/config"
	"github.com/parinpan/purwalenta/pkg/entity"
	_interface "github.com/parinpan/purwalenta/pkg/interface"
	"github.com/parinpan/purwalenta/pkg/service/request"
	"github.com/parinpan/purwalenta/pkg/service/response"
	"github.com/parinpan/purwalenta/pkg/util"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OauthService struct {
	UserRepo        _interface.UserRepository
	GoogleOauthRepo _interface.OauthRepository
}

func (service *OauthService) Exchange(ctx echo.Context, req request.OauthExchange) (response.OauthExchange, error) {
	var resp response.OauthExchange
	var providerCfg = getOauthConfig(req.Source, config.GetConfig().Oauth2)

	if nil == providerCfg {
		return resp, errors.New("can't find provider")
	}

	token, err := providerCfg.Exchange(context.TODO(), req.Code)
	if nil != err {
		return resp, errors.New("fail get token " + err.Error())
	}

	if nil == token || token.Expiry.Before(time.Now()) {
		return resp, errors.New("token is expired")
	}

	user, err := getOauthUserInfo(req.Source, service)(ctx, entity.Oauth{AccessToken: token.AccessToken})
	if nil != err {
		return resp, err
	}

	registered, err := service.UserRepo.SignUp(ctx, *user)
	if nil != err {
		return resp, err
	}

	resp.Success = registered
	resp.User.ID = user.ID
	resp.User.Username = user.Username
	resp.User.Email = user.Email
	resp.User.ProfilePicture = user.ProfilePicture
	resp.User.OauthToken = token.AccessToken
	resp.User.RefreshToken = token.RefreshToken
	resp.User.Token, _ = util.GenerateUserLoginToken(config.GetConfig(), *user)

	return resp, nil
}

func getOauthConfig(provider string, cfg config.Oauth2Config) *oauth2.Config {
	switch provider {
	case "google":
		return getGoogleOauthConfig(cfg)
	}

	return nil
}

func getOauthUserInfo(provider string, service *OauthService) func(echo.Context, entity.Oauth) (*entity.User, error) {
	switch provider {
	case "google":
		return service.GoogleOauthRepo.GetUserInfo
	}

	return nil
}

func getGoogleOauthConfig(cfg config.Oauth2Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.Google.ClientID,
		ClientSecret: cfg.Google.ClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  cfg.Google.CallbackURL,
		Scopes:       cfg.Google.Scopes,
	}
}
