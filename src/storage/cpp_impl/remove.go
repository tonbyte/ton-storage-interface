package cpp_impl

import (
	"fmt"
	"strings"
)

func (s *CppStorage) Remove(bagID string) bool {
	query := fmt.Sprintf(`"remove" "%s" "--remove-files"`, bagID)
	output, _ := s.execQuery(query)

	return strings.Contains(output, successMark)
}
