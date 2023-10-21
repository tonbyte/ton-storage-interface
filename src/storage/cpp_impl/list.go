package cpp_impl

import (
	"regexp"
	"strings"

	"tonbyte.com/ton-storage-interface/src/entity"
)

func (s *CppStorage) List() ([]entity.BagID, error) {
	query := `"list" "--hashes"`
	output, _ := s.execQuery(query)

	lines := strings.Split(output, "\n")
	if len(lines) < 2 {
		return make([]entity.BagID, 0), nil
	}

	r := regexp.MustCompile(regexpBagID)
	result := make([]entity.BagID, 0, len(lines)-2)
	for _, line := range lines {
		found := r.FindString(line)
		if found == "" {
			continue
		}

		result = append(result, entity.BagID{
			BagID: found,
		})
	}

	return result, nil
}
