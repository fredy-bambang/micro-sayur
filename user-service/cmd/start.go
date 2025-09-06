package cmd

import (
	"user-service/internal/app"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start",
	Long:  "start",
	Run: func(cmd *cobra.Command, args []string) {
		// Start the user service
		app.RunServer()
	},
}
