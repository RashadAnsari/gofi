package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"go.uber.org/automaxprocs/maxprocs"

	"github.com/RashadAnsari/gofi/internal/app/gofi/cmd"
	"github.com/RashadAnsari/gofi/pkg/log"
)

func main() {
	log.SetupLogger(log.Config{
		Level:  "debug",
		StdOut: true,
	})

	_, err := maxprocs.Set(maxprocs.Logger(logrus.Printf))
	if err != nil {
		logrus.Fatalf("failed to set GOMAXPROCS: %s", err.Error())
	}

	root := cmd.NewRootCommand()
	if root != nil {
		if err := root.Execute(); err != nil {
			os.Exit(1)
		}
	}
}
