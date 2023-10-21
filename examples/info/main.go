package main

import (
	"encoding/json"

	"tonbyte.com/ton-storage-interface/src/storage"
	"tonbyte.com/ton-storage-interface/src/storage/cpp_impl"
)

func main() {
	// tuStorage := tonutils_impl.NewTonutilsStorage(tonutils_impl.Config{
	// 	URL:   "127.0.0.1",
	// 	Port:  "38112",
	// 	Login: "admin",
	// 	Token: "123456",
	// })

	// s := storage.NewTonStorageService(tuStorage)

	orStorage := cpp_impl.NewCppStorage(cpp_impl.Config{
		CLIPath:  "/media/storage/storage-daemon/storage-daemon-cli",
		Server:   "127.0.0.1",
		CLIPort:  "5555",
		KeysPath: "/media/storage-db/cli-keys",
	})

	s := storage.NewTonStorageService(orStorage)

	res, err := s.Storage.Info("96BB34DB8EB7620574AE528E5A7FABF57B36453D8AC81A16D52EE12DC3A538BE")
	if err != nil {
		panic(err)
	}

	output, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	println(string(output))
}
