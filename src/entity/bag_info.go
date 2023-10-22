package entity

type BagInfoCommon struct {
	DownloadSpeed float32 `json:"download_speed"`
	UploadSpeed   float32 `json:"upload_speed"`
	DirName       string  `json:"dir_name"`
	Description   string  `json:"description"`
	Completed     bool    `json:"completed"`
}

type BagInfoCpp struct {
	Torrent struct {
		BagInfoCommon
		FilesCount       string `json:"files_count,omitempty"`
		TotalSize        string `json:"total_size,omitempty"`
		DownloadedSize   string `json:"downloaded_size,omitempty"`
		Hash             string `json:"hash,omitempty"`
		IncludedSize     string `json:"included_size,omitempty"`
		Flags            int32  `json:"flags,omitempty"`
		AddedAt          int32  `json:"added_at,omitempty"`
		RootDir          string `json:"root_dir,omitempty"`
		ActiveDownloaded bool   `json:"active_downloaded,omitempty"`
		ActiveUploaded   bool   `json:"active_uploaded,omitempty"`
	} `json:"torrent"`
	Files []CppStorageFile `json:"files"`
}

type BagInfoTonutils struct {
	BagInfoCommon
	Files          []TonutilsStorageFile `json:"files,omitempty"`
	FilesCount     int64                 `json:"files_count,omitempty"`
	TotalSize      int64                 `json:"size,omitempty"`
	DownloadedSize int64                 `json:"downloaded,omitempty"`
	BagPiecesNum   int64                 `json:"bag_pieces_num,omitempty"`
	HeaderLoaded   bool                  `json:"header_loaded,omitempty"`
	InfoLoaded     bool                  `json:"info_loaded,omitempty"`
	Active         bool                  `json:"active,omitempty"`
	Seeding        bool                  `json:"seeding,omitempty"`
}
