package server

import (
	"github.com/spf13/cobra"

	"github.com/vmedinskiy/highload/internal/app/server"
)

// serverCmd represents the serve command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "runs server",
	Long:  "runs server",
	RunE:  server.RunCmd,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
