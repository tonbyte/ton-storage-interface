package tonutils_impl

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *TonutilsStorage) Remove(bagID string) bool {
	url := fmt.Sprintf("http://%s:%s/api/v1/remove", s.config.URL, s.config.Port)

	body := struct {
		BagID     string `json:"bag_id"`
		WithFiles bool   `json:"with_files"`
	}{
		BagID:     bagID,
		WithFiles: true,
	}

	resp, err := s.requset(url, body, http.MethodGet)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false
	}

	var ok struct {
		Ok bool `json:"ok"`
	}
	err = json.NewDecoder(resp.Body).Decode(&ok)
	if err != nil {
		return false
	}

	return ok.Ok
}
