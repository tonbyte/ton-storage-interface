package cpp_impl

import (
	"fmt"
	"strings"
)

func (s *CppStorage) AddByBagID(bagID string, path string) (bool, error) {
	query := fmt.Sprintf(`"add-by-hash" "%s" "-d" "%s"`, bagID, path)
	output, _ := s.execQuery(query)

	if strings.Contains(output, createdMark) {
		return len(findBagID(output)) == 64, nil
	} else if strings.Contains(output, duplicateMark) {
		return len(findBagID(output)) == 64, nil
	}

	return false, ErrUnknownOutput
}
