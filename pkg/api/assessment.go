package api

import (
	_interface "github.com/purwalenta/purwalenta/pkg/interface"
	"github.com/purwalenta/purwalenta/pkg/repository/static"
	servicePkg "github.com/purwalenta/purwalenta/pkg/service"
)

var (
	assessmentInstances = make(map[Type]*AssessmentAPI)
)

type AssessmentAPI struct {
	Service _interface.AssessmentService
}

func NewAssessmentAPI(apiType Type) *AssessmentAPI {
	if _, exists := assessmentInstances[apiType]; exists {
		return assessmentInstances[apiType]
	}

	switch apiType {
	case DefaultAssessmentAPIFlag:
		return newDefaultAssessmentAPI()
	}

	return newDefaultAssessmentAPI()
}

func newDefaultAssessmentAPI() *AssessmentAPI {
	service := new(servicePkg.AssessmentService)
	service.Repo = new(static.AssessmentRepository)

	assessmentInstances[DefaultAssessmentAPIFlag] = &AssessmentAPI{
		Service: service,
	}

	return assessmentInstances[DefaultAssessmentAPIFlag]
}
