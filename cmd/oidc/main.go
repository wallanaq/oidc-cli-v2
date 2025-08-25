package main

import (
	"github.com/spf13/cobra"
	"github.com/wallanaq/oidc-cli/v2/cmd/oidc/root"
)

func main() {

	rootCmd := root.NewRootCmd()

	if err := rootCmd.Execute(); err != nil {
		cobra.CheckErr(err)
	}

}
