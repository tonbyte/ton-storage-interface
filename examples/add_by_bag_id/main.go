package main

import (
	"github.com/tonbyte/ton-storage-interface/src/storage"
	"github.com/tonbyte/ton-storage-interface/src/storage/cpp_impl"
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

	// res, err := s.Storage.AddByBagID("", "/home/storage/utilsstorage/tonutils-storage-db/downloads")
	res, err := s.Storage.AddByBagID("F0644CE51A3584590BA29E20FB9B84E9BD94007345D94D6EE112C6179A09476A", "/home/storage")
	if err != nil {
		panic(err)
	}

	println(res)
}
