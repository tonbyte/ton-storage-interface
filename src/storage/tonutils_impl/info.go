package tonutils_impl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tonbyte/ton-storage-interface/src/entity"
)

func (s *TonutilsStorage) Info(bagID string) (interface{}, error) {
	url := fmt.Sprintf("http://%s:%s/api/v1/details?bag_id=%s", s.config.URL, s.config.Port, bagID)

	resp, err := s.requset(url, nil, http.MethodGet)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", resp.Status)
	}

	var info entity.BagInfoTonutils
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil, err
	}

	return info, nil
}
