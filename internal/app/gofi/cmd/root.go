package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/RashadAnsari/gofi/internal/app/gofi/cmd/server"
	"github.com/RashadAnsari/gofi/internal/app/gofi/cmd/version"
	"github.com/RashadAnsari/gofi/internal/app/gofi/config"
	"github.com/RashadAnsari/gofi/pkg/log"
)

func NewRootCommand() *cobra.Command {
	var root = &cobra.Command{
		Use:   "gofi",
		Short: "Gofi Service.",
	}

	root.CompletionOptions.DisableDefaultCmd = true

	cfg, err := config.Init()
	if err != nil {
		root.PrintErrf("Failed to initialize configuration: %s\n", err.Error())
		os.Exit(1)
	}

	log.SetupLogger(cfg.Logger)

	version.Register(root)
	server.Register(root, *cfg)

	return root
}
