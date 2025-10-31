package types

type RootResponse struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
	Uptime  string `json:"uptime"`
	Version string `json:"version"`
}

type LinkResponse struct {
	Ok           bool   `json:"ok"`
	MessageID    int    `json:"message_id,omitempty"`
	FileName     string `json:"file_name,omitempty"`
	FileSize     int64  `json:"file_size,omitempty"`
	MimeType     string `json:"mime_type,omitempty"`
	DownloadLink string `json:"download_link,omitempty"`
	Error        string `json:"error,omitempty"`
}
