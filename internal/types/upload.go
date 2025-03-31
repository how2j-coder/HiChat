package types

type CreatUploadMultiReq struct {
	Type             string `json:"type"`
	MimeType         string `json:"mime_type"`
	Size             int64  `json:"size"`
	MD5              string `json:"md5"`
	Key              string `json:"key"`
	Storage          string `json:"storage"`
	OriginalFileName string `json:"original_file_name"`
	RelativeFile     string `json:"relative_file"`
	URL              string `json:"url"`
	Status           int8   `json:"status"`
	Error            string `json:"error,omitempty"`
}
