package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CurrentVersion string = ""

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of epub2pdf",
	Long:  `Print the current version number of epub2pdf`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(CurrentVersion)
	},
}
