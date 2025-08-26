package updatecheck

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/wallanaq/oidc-cli/v2/internal/version"
)

func NewUpdateCheckCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "update-check",
		Short: "Check for updates",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithTimeout(cmd.Context(), 3*time.Second)
			defer cancel()

			updateInfo, err := version.CheckForUpdate(ctx)
			if err != nil {
				return fmt.Errorf("could not check for updates: %w", err)
			}

			if updateInfo.UpdateAvailable {
				cmd.Println("An update is available")
				cmd.Printf("Current: %s\n", updateInfo.CurrentVersion)
				cmd.Printf("Latest : %s\n", updateInfo.LatestVersion)
				return nil
			}

			cmd.Println("No updates available")

			return nil
		},
	}

}
