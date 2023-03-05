package server

import (
	"github.com/spf13/cobra"

	"github.com/vmedinskiy/highload/internal/app/migrate"
)

// serverCmd represents the serve command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "runs db migration",
	Long:  "runs db migration",
	RunE:  migrate.RunCmd,
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
