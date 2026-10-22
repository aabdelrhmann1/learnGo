package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rssagg",
	Short: "An application for user feeds",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("specify a command, for example: rssagg feed")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
