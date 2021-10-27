package logic

import (
	"encoding/json"
	"github.com/klovercloud-ci/config"
	"github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/service"
	"log"
)

type processEventService struct {
	httpPublisher service.HttpClient
}

func (e processEventService) Listen(subject v1.Subject) {
		event:=v1.ProcessEvent{
			ProcessId: subject.Pipeline.ProcessId,
			Data:      subject.EventData,
		}
		header:=make(map[string]string)
		header["Content-Type"]="application/json"
		header["token"]=config.Token
		b, err := json.Marshal(event)
		if err!=nil{
			log.Println(err.Error())
			return
		}
		e.httpPublisher.Post(config.EventStoreUrl+"/processes_events",header,b)
}

func NewProcessEventService(httpPublisher service.HttpClient) service.Observer {
	return &processEventService{
		httpPublisher: httpPublisher,
	}
}
