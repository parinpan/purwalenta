package static

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo"
	"github.com/purwalenta/purwalenta/pkg/config"
	"github.com/purwalenta/purwalenta/pkg/entity"
)

type AssessmentRepository struct {
}

func (repo *AssessmentRepository) GetPersonalityQuestions(ctx echo.Context) (entity.PersonalityAssessmentQuestion, error) {
	var result entity.PersonalityAssessmentQuestion

	fmt.Println(config.GetConfig().Assessment.PersonalityQuestionsJSON)

	if err := json.Unmarshal([]byte(config.GetConfig().Assessment.PersonalityQuestionsJSON), &result); nil != err {
		fmt.Println(err)
		return result, err
	}

	fmt.Println(result)

	return result, nil
}
