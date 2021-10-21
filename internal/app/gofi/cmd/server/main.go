package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/RashadAnsari/gofi/pkg/s3"

	"github.com/sirupsen/logrus"

	"github.com/RashadAnsari/gofi/internal/app/gofi/config"
	"github.com/RashadAnsari/gofi/internal/app/gofi/router"

	"github.com/spf13/cobra"
)

func main(cfg config.Config) {
	ctx, cancl := context.WithCancel(context.Background())

	defer cancl()

	echoRouter := router.New(cfg)

	s3Client, err := s3.Create(ctx, cfg.S3, config.DefaultS3Bucket)
	if err != nil {
		logrus.Fatalf("failed to create s3 client: %s", err.Error())
	}

	registerRouters(serviceConfig{
		router:   echoRouter,
		s3Client: s3Client,
	})

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := echoRouter.Start(cfg.Server.Address); err != nil {
			logrus.Fatalf("failed to start the server: %s", err.Error())
		}
	}()

	logrus.Infof("starting the server: http://%s", cfg.Server.Address)

	s := <-sig

	logrus.Infof("signal %s received", s.String())

	echoCtx, echoCancel := context.WithTimeout(ctx, cfg.Server.GracefulTimeout)

	defer echoCancel()

	if err := echoRouter.Shutdown(echoCtx); err != nil {
		logrus.Errorf("failed to shutdown the server: %s", err.Error())
	}
}

func Register(root *cobra.Command, cfg config.Config) {
	root.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "Run the Gofi Service",
			Run: func(cmd *cobra.Command, args []string) {
				main(cfg)
			},
		},
	)
}
