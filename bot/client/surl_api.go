package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/Migan178/surl-bot/configs"
)

type SURL struct {
	Client *http.Client
}

var instance *SURL

func GetClient() *SURL {
	if instance == nil {
		instance = &SURL{&http.Client{}}
	}

	return instance
}

func (s *SURL) GetInformation(link string) (*URL, error) {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Get(fmt.Sprintf("%s/links%s", configs.GetConfig().SURL.API, parsedURL.Path))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	dataBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var data errResponse

		err = json.Unmarshal(dataBytes, &data)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("http status code is not 200. message: %s", data.Message)
	}

	var data URL

	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *SURL) Create(link string) (*URL, error) {
	dataString := fmt.Sprintf("{\"redirect_url\":\"%s\"}", link)

	buf := bytes.NewBuffer([]byte(dataString))

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/links", configs.GetConfig().SURL.API), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	dataBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		var data errResponse

		err = json.Unmarshal(dataBytes, &data)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("http status code is not 201. message: %s", data.Message)
	}

	var data URL

	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
