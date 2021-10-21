package file

import (
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/awserr"

	"github.com/RashadAnsari/gofi/internal/app/gofi/config"
	jsonTime "github.com/RashadAnsari/gofi/pkg/time"

	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/RashadAnsari/gofi/internal/app/gofi/handler"
)

const (
	accessHashMeta     = "Access-Hash"
	presignedURLExpiry = 5 * time.Minute
)

type Handler struct {
	S3Client *s3.S3
}

// swagger:route GET /v1/file/upload File GetFileUploadURL
//
// Get File Upload URL
//
// By using this endpoint, you will get a URL for uploading your file (For less than 5 MB file size).
//
//	Responses:
//		200: GetFileUploadURLResponse
//		400: Error
//		500: Error
func (h Handler) GetFileUploadURL(c echo.Context) error {
	ctx := c.Request().Context()

	log := handler.LogEntry(ctx, "file", "GetFileUploadURL")

	req := GetFileUploadURLRequest{}

	if err := (&echo.DefaultBinder{}).BindHeaders(c, &req); err != nil {
		log.Errorf("failed to bind request: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, handler.ErrInvalidHeaderSyntax.Error())
	}

	if err := req.Validate(); err != nil {
		log.Errorf("failed to validate request: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	fileID := uuid.New().String()
	accessHash := uuid.New().String()

	input := &s3.PutObjectInput{
		Key:           &fileID,
		Bucket:        aws.String(config.DefaultS3Bucket),
		ContentType:   &req.ContentType,
		ContentLength: &req.ContentLength,
		Metadata: map[string]*string{
			accessHashMeta: &accessHash,
		},
	}

	putReq, _ := h.S3Client.PutObjectRequest(input)

	url, headers, err := putReq.PresignRequest(presignedURLExpiry)
	if err != nil {
		log.Errorf("failed to get pre-signed PUT URL from S3: %s", err.Error())
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, GetFileUploadURLResponse{
		FileID:        fileID,
		AccessHash:    accessHash,
		UploadURL:     url,
		ExpireAt:      jsonTime.JSONTime(time.Now().Add(presignedURLExpiry)),
		UploadHeaders: headers,
	})
}

// swagger:route GET /v1/file/download File GetFileDownloadURL
//
// Get File Download URL
//
// By using this endpoint, you will get a URL for downloading your file.
//
//	Responses:
//		200: GetFileDownloadURLResponse
//		400: Error
//		403: Error
//		500: Error
func (h Handler) GetFileDownloadURL(c echo.Context) error {
	ctx := c.Request().Context()

	log := handler.LogEntry(ctx, "file", "GetFileDownloadURL")

	req := GetFileDownloadURLRequest{}

	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &req); err != nil {
		log.Errorf("failed to bind request: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, handler.ErrInvalidQuerySyntax.Error())
	}

	if err := req.Validate(); err != nil {
		log.Errorf("failed to validate request: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	headObjectInput := &s3.HeadObjectInput{
		Key:    &req.FileID,
		Bucket: aws.String(config.DefaultS3Bucket),
	}

	headObject, err := h.S3Client.HeadObjectWithContext(ctx, headObjectInput)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Message() == http.StatusText(http.StatusNotFound) {
			return echo.NewHTTPError(http.StatusForbidden, "unable to access to this resource")
		}

		log.Errorf("failed to get head object from S3: %s", err.Error())

		return echo.ErrInternalServerError
	}

	accessHash, ok := headObject.Metadata[accessHashMeta]
	if !ok || (ok && accessHash == nil) || (ok && accessHash != nil && *accessHash != req.AccessHash) {
		return echo.NewHTTPError(http.StatusForbidden, "unable to access to this resource")
	}

	getObjectInput := &s3.GetObjectInput{
		Key:    &req.FileID,
		Bucket: aws.String(config.DefaultS3Bucket),
	}

	getReq, _ := h.S3Client.GetObjectRequest(getObjectInput)

	url, err := getReq.Presign(presignedURLExpiry)
	if err != nil {
		log.Errorf("failed to get pre-signed Get URL from S3: %s", err.Error())
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, GetFileDownloadURLResponse{
		DownloadURL:   url,
		ExpireAt:      jsonTime.JSONTime(time.Now().Add(presignedURLExpiry)),
		ContentType:   headObject.ContentType,
		ContentLength: headObject.ContentLength,
	})
}
