package file

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

const (
	oneKiloByte  = 1 * 1024
	fileMegaByte = 5 * 1024 * 1024
)

type GetFileUploadURLRequest struct {
	ContentType   string `json:"Content-Type" header:"Content-Type"`
	ContentLength int64  `json:"Content-Length" header:"Content-Length"`
}

func (g GetFileUploadURLRequest) Validate() error {
	return validation.ValidateStruct(&g,
		validation.Field(&g.ContentType,
			validation.Required.Error("the Content-Type HTTP request header is required"),
			validation.In("image/png", "image/jpeg").Error("the Content-Type HTTP request header is not allowed"),
		),
		validation.Field(&g.ContentLength,
			validation.Required.Error("the Content-Length HTTP request header is required"),
			validation.Min(oneKiloByte).Error("the file size is less than the allowed minimum of 1 KB"),
			validation.Max(fileMegaByte).Error("the file size is bigger than the allowed maximum of 5 MB"),
		),
	)
}

// swagger:parameters GetFileUploadURL
type GetFileUploadURLRequestWrapper struct {
	// required: true
	// in: header
	ContentType string `json:"Content-Type"`
	// required: true
	// in: header
	ContentLength int64 `json:"Content-Length"`
}

type GetFileDownloadURLRequest struct {
	FileID     string `json:"fileId" query:"fileId"`
	AccessHash string `json:"accessHash" query:"accessHash"`
}

func (g GetFileDownloadURLRequest) Validate() error {
	return validation.ValidateStruct(&g,
		validation.Field(&g.FileID,
			validation.Required.Error("the fileId HTTP request query param is required"),
			is.UUID.Error("the fileId HTTP request query param is invalid"),
		),
		validation.Field(&g.AccessHash,
			validation.Required.Error("the accessHash HTTP request query param is required"),
			is.UUID.Error("the fileId HTTP request query param is invalid"),
		),
	)
}

// swagger:parameters GetFileDownloadURL
type GetFileDownloadURLRequestWrapper struct {
	// required: true
	// in: query
	FileID string `json:"fileId"`
	// required: true
	// in: query
	AccessHash string `json:"accessHash"`
}
