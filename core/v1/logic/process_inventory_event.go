package logic

import (
	"encoding/json"
	"github.com/klovercloud-ci/config"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/service"
	"log"
)

type processInventoryEventService struct {
	httpClient service.HttpClient
}

func (p processInventoryEventService) Listen(subject v1.Subject) {
	if subject.App.CompanyId == "" {
		return
	}
	url := config.EventStoreUrl + "/processes"

	header := make(map[string]string)
	header["token"]=config.Token
	header["Content-Type"] = "application/json"

	process := v1.ProcessInventoryEvent{
		ProcessId:    subject.Pipeline.ProcessId,
		CompanyId:    subject.App.CompanyId,
		AppId:        subject.App.AppId,
		RepositoryId: subject.App.RepositoryId,
		Data:         nil,
	}
	b, err := json.Marshal(process)
	if err != nil {
		log.Println(err.Error())
		return
	}
	go p.httpClient.Post(url, header, b)
}

func NewProcessInventoryEventService(httpPublisher service.HttpClient) service.Observer {
	return &processInventoryEventService{
		httpClient: httpPublisher,
	}
}
