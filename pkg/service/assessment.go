package service

import (
	"github.com/labstack/echo"
	_interface "github.com/purwalenta/purwalenta/pkg/interface"
	"github.com/purwalenta/purwalenta/pkg/service/request"
	"github.com/purwalenta/purwalenta/pkg/service/response"
)

type AssessmentService struct {
	repo _interface.AssessmentRepository
}

func (service *AssessmentService) FindPersonalityQuestions(ctx echo.Context, req request.FindPersonalityQuestions) (response.FindPersonalityQuestions, error) {
	return response.FindPersonalityQuestions{}, nil
}
