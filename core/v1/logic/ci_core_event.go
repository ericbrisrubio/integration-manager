package logic

import (
	"encoding/json"
	"github.com/klovercloud-ci/config"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/service"
	"log"
)

type ciCoreEventService struct {
	httpClient service.HttpClient
}

func (a ciCoreEventService) Listen(subject v1.Subject) {
	if subject.EventData != nil && len(subject.EventData) > 0 {
		if subject.EventData["trigger"] == false {
			return
		}
	}
	if subject.CoreRequestQueryParam == nil {
		return
	}
	url := config.KlovercloudCiCoreUrl + "/pipelines?url=" + subject.CoreRequestQueryParam["url"] + "&revision=" + subject.CoreRequestQueryParam["revision"] + "&purging=" + subject.CoreRequestQueryParam["purging"]

	header := make(map[string]string)
	header["token"] = config.Token
	header["Content-Type"] = "application/json"

	b, err := json.Marshal(subject.Pipeline)
	if err != nil {
		log.Println(err.Error())
		return
	}
	go a.httpClient.Post(url, header, b)
}

// NewCiCoreEventService returns Observer type service
func NewCiCoreEventService(httpPublisher service.HttpClient) service.Observer {
	return &ciCoreEventService{
		httpClient: httpPublisher,
	}
}
