package v1

import (
	"github.com/klovercloud-ci-cd/integration-manager/api/common"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/api"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/labstack/echo/v4"
	"log"
)

type agentApi struct {
	agentService service.Agent
}

// Store... Store agent
// @Summary Store agent
// @Description Stores agent
// @Tags Agent
// @Produce json
// @Param data body v1.Agent true "Agent Name"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/agent [POST]
func (a agentApi) Store(context echo.Context) error {
	formData := v1.Agent{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	err := a.agentService.Store(formData)
	if err != nil {
		log.Println("[Error]:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Operation Failed!")
	}
	return common.GenerateSuccessResponse(context, nil,
		nil, "Operation Successful")
}

// NewCompanyApi returns Agent type api
func NewAgentApi(agentService service.Agent) api.Agent {
	return &agentApi{
		agentService: agentService,
	}
}
