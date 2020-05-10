package service

import (
	"github.com/labstack/echo"
	_interface "github.com/parinpan/purwalenta/pkg/interface"
	"github.com/parinpan/purwalenta/pkg/service/request"
	"github.com/parinpan/purwalenta/pkg/service/response"
)

type AssessmentService struct {
	Repo _interface.AssessmentRepository
}

func (service *AssessmentService) FindPersonalityQuestions(ctx echo.Context, req request.FindPersonalityQuestions) (response.PersonalityQuestion, error) {
	var questions, _ = service.Repo.GetPersonalityQuestions(ctx)
	return response.PersonalityQuestion(questions), nil
}
