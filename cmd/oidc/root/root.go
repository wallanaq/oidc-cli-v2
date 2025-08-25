package root

import (
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/wallanaq/oidc-cli/v2/cmd/oidc/version"
)

func NewRootCmd() *cobra.Command {

	var debugMode bool

	rootCmd := &cobra.Command{
		Use:           "oidc",
		Short:         "OpenID Connect command-line tool",
		SilenceErrors: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if debugMode {
				slog.SetLogLoggerLevel(slog.LevelDebug)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.AddCommand(version.NewVersionCmd())

	rootCmd.PersistentFlags().BoolVarP(&debugMode, "debug", "d", false, "Enable debug mode")

	return rootCmd

}
