package apicalls

import (
	"encoding/json"

	_ "github.com/go-resty/resty"
	"github.com/labstack/echo"
	"github.com/purwalenta/purwalenta/pkg/entity"
	"github.com/purwalenta/purwalenta/pkg/repository/apicalls/response"
)

const (
	oauthGoogleURL = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
)

type GoogleOauthRepository struct {
}

func (repo *GoogleOauthRepository) GetUserInfo(ctx echo.Context, oauth entity.Oauth) (*entity.User, error) {
	var user = new(entity.User)

	apiResp, err := getClient().R().EnableTrace().Get(oauthGoogleURL + oauth.AccessToken)
	if nil != err {
		return user, nil
	}

	// reserve a struct that will hold the api response data
	var resp = new(response.GoogleAuthUserInfo)

	if err := json.Unmarshal(apiResp.Body(), &resp); nil != err {
		return user, err
	}

	user.Username = resp.ID
	user.Email = resp.Email
	user.FullName = resp.Name
	user.ProfilePicture = resp.Picture

	return user, nil
}
