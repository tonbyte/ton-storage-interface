package cpp_impl

import (
	"fmt"
	"strings"
)

func (s *CppStorage) Create(path string) (string, bool, error) {
	query := fmt.Sprintf(`"create" "--" "%s"`, path)
	output, _ := s.execQuery(query)

	if strings.Contains(output, createdMark) {
		return findBagID(output), false, nil
	} else if strings.Contains(output, duplicateMark) {
		return findBagID(output), true, nil
	}

	return "", false, ErrUnknownOutput
}
