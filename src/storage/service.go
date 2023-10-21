package storage

import (
	"fmt"

	"tonbyte.com/ton-storage-interface/src/entity"
)

type TonStorage interface {
	Info(string) (interface{}, error)
	List() ([]entity.BagID, error)
	Create(string) (string, bool, error)
	AddByBagID(string, string) (bool, error)
	Remove(string) bool
}

var (
	ErrCanNotParseResponse = fmt.Errorf("can not parse response")
)

type TonStorageService struct {
	Storage TonStorage
}

func NewTonStorageService(storage TonStorage) *TonStorageService {
	return &TonStorageService{Storage: storage}
}
