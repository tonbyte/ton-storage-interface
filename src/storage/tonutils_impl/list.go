package tonutils_impl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"tonbyte.com/ton-storage-interface/src/entity"
)

func (s *TonutilsStorage) List() ([]entity.BagID, error) {
	url := fmt.Sprintf("http://%s:%s/api/v1/list", s.config.URL, s.config.Port)

	resp, err := s.requset(url, nil, http.MethodGet)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", resp.Status)
	}

	var info = struct {
		Bags []entity.BagID `json:"bags"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil, err
	}

	return info.Bags, nil
}
