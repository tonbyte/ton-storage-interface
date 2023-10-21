package cpp_impl

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tonbyte/ton-storage-interface/src/entity"
	"github.com/tonbyte/ton-storage-interface/src/storage"
)

func (s *CppStorage) Info(bagID string) (interface{}, error) {
	query := fmt.Sprintf(`"get" "--json" "%s"`, bagID)
	output, _ := s.execQuery(query)

	if !strings.Contains(output, `"@type"`) {
		return nil, storage.ErrCanNotParseResponse
	}

	start := strings.Index(output, "{")
	end := strings.LastIndex(output, "}")

	if start == -1 || end == -1 {
		return nil, storage.ErrCanNotParseResponse
	}

	output = output[start : end+1]
	result := entity.BagInfoCpp{}
	err := json.Unmarshal([]byte(output), &result)
	if err != nil {
		return nil, storage.ErrCanNotParseResponse
	}

	return result, nil
}
