package entity

type File struct {
	Name string `json:"name"`
}

type CppStorageFile struct {
	File
	Size string `json:"size"`
}

type TonutilsStorageFile struct {
	File
	Size int64 `json:"size"`
}
