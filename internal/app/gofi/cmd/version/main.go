package version

import (
	"github.com/spf13/cobra"

	"github.com/RashadAnsari/gofi/pkg/version"
)

func Register(root *cobra.Command) {
	root.AddCommand(
		&cobra.Command{
			Use:   "version",
			Short: "Print the version of the Gofi Service",
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Println(version.String())
			},
		},
	)
}
