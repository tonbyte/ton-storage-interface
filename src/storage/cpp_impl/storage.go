package cpp_impl

import (
	"fmt"
	"os/exec"
	"regexp"
)

type Config struct {
	CLIPath  string
	Server   string
	CLIPort  string
	KeysPath string
}

var (
	ErrUnknownOutput = fmt.Errorf("unknown output")
)

const (
	successMark   = "Success"
	createdMark   = "Bag created"
	duplicateMark = "duplicate hash "
	regexpBagID   = "([0-9A-F]{64})"
)

type CppStorage struct {
	Config Config
}

func NewCppStorage(config Config) *CppStorage {
	return &CppStorage{
		Config: config,
	}
}

func (s *CppStorage) execQuery(query string) (string, error) {
	command := exec.Command(s.Config.CLIPath,
		`-I`, s.Config.Server+`:`+s.Config.CLIPort,
		`-k`, s.Config.KeysPath+`/client`,
		`-p`, s.Config.KeysPath+`/server.pub`,
		`-c`, query)

	output, err := command.CombinedOutput()
	return string(output), err
}

func findBagID(output string) string {
	r := regexp.MustCompile(regexpBagID)
	result := r.FindString(output)

	return result
}
