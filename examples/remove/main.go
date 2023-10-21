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

	res := s.Storage.Remove("4f1309837dea51ce7c0d6e4232fe8b7cab9ad1ab41cb4e0fa193143ac68d7908")

	output, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	println(string(output))
}
