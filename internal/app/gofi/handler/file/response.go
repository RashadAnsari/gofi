// nolint: lll, godot
package file

import "github.com/RashadAnsari/gofi/pkg/time"

// swagger:model
type GetFileUploadURLResponse struct {
	// example: 46ef48d4-224e-44a6-ab2f-0e71efddd248
	FileID string `json:"file_id"`
	// example: cd001b1c-c482-4bb8-a51b-6fcd0236487d
	AccessHash string `json:"access_hash"`
	// example: http://127.0.0.1:9000/gofi/46ef48d4-224e-44a6-ab2f-0e71efddd248?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=access-key%2F20211019%2Feu-east-1%2Fs3%2Faws4_request&X-Amz-Date=20211019T104626Z&X-Amz-Expires=300&X-Amz-SignedHeaders=content-length%3Bcontent-type%3Bhost%3Bx-amz-meta-access-hash&X-Amz-Signature=01cb5e304b2a3341a14bca2fd3fb074b369c5e8d49221d5511e312859e1bcd9a
	UploadURL string `json:"upload_url"`
	// RFC 3339 time format.
	// example: 2021-10-19T13:42:35+02:00
	ExpireAt time.JSONTime `json:"expire_at"`
	// example: {"content-length": ["1977"], "content-type": ["image/png"], "x-amz-meta-access-hash": ["cd001b1c-c482-4bb8-a51b-6fcd0236487d"]}
	UploadHeaders map[string][]string `json:"upload_headers"`
}

// Get file upload URL response.
// swagger:response GetFileUploadURLResponse
type GetFileUploadURLResponseWrapper struct {
	// in: body
	_ GetFileUploadURLResponse
}

// swagger:model
type GetFileDownloadURLResponse struct {
	// example: http://127.0.0.1:9000/gofi/406ab97e-eb4d-4697-9d37-9b890d2af7b4?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=access-key%2F20211019%2Feu-east-1%2Fs3%2Faws4_request&X-Amz-Date=20211019T114140Z&X-Amz-Expires=300&X-Amz-SignedHeaders=host&X-Amz-Signature=5c4de9845ab5ed7726fbe81b913f1564f355ea303e324affa53d8e5e9549ea25
	DownloadURL string `json:"download_url"`
	// RFC 3339 time format.
	// example: 2021-10-19T13:42:35+02:00
	ExpireAt time.JSONTime `json:"expire_at"`
	// example: image/png
	ContentType *string `json:"content_type,omitempty"`
	// example: 1977
	ContentLength *int64 `json:"content_length,omitempty"`
}

// Get file download URL response.
// swagger:response GetFileDownloadURLResponse
type GetFileDownloadURLResponseWrapper struct {
	// in: body
	_ GetFileDownloadURLResponse
}
