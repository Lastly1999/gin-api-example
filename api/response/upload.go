package response

type UploadFileRes struct {
	Url  string `json:"url"`
	Size int64  `json:"size"`
	Name string `json:"name"`
}
