package handler

import (
	"github.com/labstack/echo"
)

type assessmentHandler struct {
}

func NewAssessmentHandler() *assessmentHandler {
	return &assessmentHandler{}
}

func (handler *assessmentHandler) Personality(ctx echo.Context) error {
	return nil
}
