package tonutils_impl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"tonbyte.com/ton-storage-interface/src/entity"
)

func (s *TonutilsStorage) Create(path string) (string, bool, error) {
	url := fmt.Sprintf("http://%s:%s/api/v1/create", s.config.URL, s.config.Port)

	body := struct {
		Path string `json:"path"`
	}{
		Path: path,
	}
	resp, err := s.requset(url, body, http.MethodPost)
	if err != nil {
		return "", false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", false, fmt.Errorf("%s", resp.Status)
	}

	var result entity.BagID
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", false, err
	}

	return result.BagID, false, nil
}
