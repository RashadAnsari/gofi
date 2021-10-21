package handler

import (
	"context"

	"github.com/sirupsen/logrus"
)

func LogEntry(ctx context.Context, handler, method string) *logrus.Entry {
	return logrus.WithContext(ctx).WithFields(logrus.Fields{
		"handler": handler,
		"method":  method,
	})
}
