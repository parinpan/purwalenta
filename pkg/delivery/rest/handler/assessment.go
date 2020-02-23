package handler

import (
	"github.com/labstack/echo"
	apiPkg "github.com/purwalenta/purwalenta/pkg/api"
	"github.com/purwalenta/purwalenta/pkg/delivery/rest/response"
	"github.com/purwalenta/purwalenta/pkg/service/request"
)

type assessmentHandler struct {
	api *apiPkg.AssessmentAPI
}

func NewAssessmentHandler() *assessmentHandler {
	return &assessmentHandler{
		api: apiPkg.NewAssessmentAPI(apiPkg.DefaultAssessmentAPIFlag),
	}
}

func (handler *assessmentHandler) FindPersonalityQuestions(ctx echo.Context) error {
	var req = new(request.FindPersonalityQuestions)
	var resp = response.New()
	var err error

	if resp.Data, err = handler.api.Service.FindPersonalityQuestions(ctx, *req); nil != err {
		return resp.ToJSON(ctx)
	}

	return resp.ToJSON(ctx)
}
