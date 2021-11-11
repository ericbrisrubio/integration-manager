package logic

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/integration-manager/config"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"log"
)

type logEventService struct {
	httpPublisher service.HttpClient
}

func (e logEventService) Listen(subject v1.Subject) {
	data := v1.LogEvent{
		ProcessId: subject.Pipeline.ProcessId,
		Log:       subject.Log,
		Step:      subject.Step,
	}
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["token"] = config.Token
	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
		return
	}
	e.httpPublisher.Post(config.EventStoreUrl+"/logs", header, b)
}

// NewLogEventService returns Observer type service
func NewLogEventService(httpPublisher service.HttpClient) service.Observer {
	return &logEventService{
		httpPublisher: httpPublisher,
	}
}
