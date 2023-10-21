package tonutils_impl

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Config struct {
	URL   string
	Port  string
	Login string
	Token string
}

type TonutilsStorage struct {
	config Config
}

func NewTonutilsStorage(config Config) *TonutilsStorage {
	return &TonutilsStorage{
		config: config,
	}
}

func (s *TonutilsStorage) requset(url string, body interface{}, method string) (*http.Response, error) {
	var req *http.Request
	var err error
	var requestBody *bytes.Buffer = nil
	if body != nil {
		var rb []byte
		rb, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}

		requestBody = bytes.NewBuffer(rb)
		req, err = http.NewRequest(method, url, requestBody)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(s.config.Login, s.config.Token)

	return http.DefaultClient.Do(req)
}
