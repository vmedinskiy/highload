package server

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

// serverCmd represents the serve command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "prints version information",
	Long:  "prints version information",
	Run: func(cmd *cobra.Command, _ []string) {
		fmt.Printf("Build version: %s\nBuild date: %s\nBuild commit: %s\n", buildVersion, buildDate, buildCommit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
