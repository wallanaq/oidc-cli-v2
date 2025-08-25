package version

import (
	"github.com/spf13/cobra"
	"github.com/wallanaq/oidc-cli/v2/internal/version"
)

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version of oidc-cli",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println(version.String())
		},
	}
}
