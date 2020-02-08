package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type assessmentHandler struct {
}

func NewAssessmentHandler() *assessmentHandler {
	return &assessmentHandler{}
}

func (handler *assessmentHandler) Index(ctx echo.Context) error {
	return ctx.JSON(http.StatusAccepted, map[string]interface{}{"ok": true})
}
