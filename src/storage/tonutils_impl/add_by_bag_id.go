package tonutils_impl

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *TonutilsStorage) AddByBagID(bagID string, path string) (bool, error) {
	url := fmt.Sprintf("http://%s:%s/api/v1/add", s.config.URL, s.config.Port)

	body := struct {
		BagID       string `json:"bag_id"`
		Path        string `json:"path"`
		DownloadAll bool   `json:"download_all"`
	}{
		BagID:       bagID,
		Path:        path,
		DownloadAll: true,
	}
	resp, err := s.requset(url, body, http.MethodPost)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("%s", resp.Status)
	}

	var ok = struct {
		Ok bool `json:"ok"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&ok)
	if err != nil {
		return false, err
	}

	return ok.Ok, nil
}
