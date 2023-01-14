package commands

import (
	"fmt"

	"github.com/kironono/weather-forecast-lamp/internal/model"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version: %s-%s\n", model.Version, model.Revision)
	},
}
