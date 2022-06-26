package logic

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/integration-manager/config"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"log"
)

type processEventService struct {
	httpPublisher service.HttpClient
}

func (e processEventService) Listen(subject v1.Subject) {
	if subject.Pipeline.ProcessId==""{
		return
	}
	event := v1.ProcessEvent{
		ProcessId: subject.Pipeline.ProcessId,
		Data:      subject.EventData,
	}
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["token"] = config.Token
	b, err := json.Marshal(event)
	if err != nil {
		log.Println(err.Error())
		return
	}
	e.httpPublisher.Post(config.EventStoreUrl+"/processes_events", header, b)
}

// NewProcessEventService returns Observer type service
func NewProcessEventService(httpPublisher service.HttpClient) service.Observer {
	return &processEventService{
		httpPublisher: httpPublisher,
	}
}
