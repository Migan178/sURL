package client

import (
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

	resp, err := s.Client.Get(fmt.Sprintf("%s/links%s", configs.GetConfig().SURL.APIURL, parsedURL.Path))
	if err != nil {
		return nil, err
	}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var data errResponse

		err = json.Unmarshal(buf, &data)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("http status code is not 200. message: %s", data.Message)
	}

	var data URL

	err = json.Unmarshal(buf, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
