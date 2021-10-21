package server

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo/v4"

	"github.com/RashadAnsari/gofi/internal/app/gofi/handler/file"

	"github.com/RashadAnsari/gofi/internal/app/gofi/handler/health"
)

const (
	v1Prefix = "/v1"
)

type serviceConfig struct {
	router   *echo.Echo
	s3Client *s3.S3
}

func registerRouters(cfg serviceConfig) {
	healthHandler := health.Handler{}
	fileHandler := file.Handler{S3Client: cfg.s3Client}

	cfg.router.GET("/healthz", healthHandler.Healthz)

	v1 := cfg.router.Group(v1Prefix)
	{
		v1.GET("/file/upload", fileHandler.GetFileUploadURL)
		v1.GET("/file/download", fileHandler.GetFileDownloadURL)
	}
}
