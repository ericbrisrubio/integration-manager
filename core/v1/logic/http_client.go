package logic

import (
	"bytes"
	"errors"
	"github.com/klovercloud-ci/core/v1/service"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type httpClientService struct {
}

func (h httpClientService) Delete(url string, header map[string]string) error {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("[ERROR] Failed communicate :", err.Error())
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return err
	}
	return nil
}

func (h httpClientService) Get(url string, header map[string]string) (error, []byte) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	for k, v := range header {
		req.Header.Set(k, v)
	}
	if err != nil {
		log.Println(err.Error())
		return err, nil
	}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return err, nil
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		jsonDataFromHttp, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err.Error())
			return err, nil
		}
		return nil, jsonDataFromHttp
	} else {
		return errors.New("Status: " + res.Status + ", code: " + strconv.Itoa(res.StatusCode)), nil
	}

}

func (h httpClientService) Post(url string, header map[string]string, body []byte) (error, []byte) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	for k, v := range header {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("[ERROR] Failed communicate :", err.Error())
		return err, nil
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return err, nil
	}
	if err != nil {
		return err, nil
	}
	return nil, body
}

func NewHttpClientService() service.HttpClient {
	return &httpClientService{}
}
