package commands

import (
	"github.com/kironono/weather-forecast-lamp/internal/device"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listDevicesCmd)
}

var listDevicesCmd = &cobra.Command{
	Use: "list-devices",
	Run: func(cmd *cobra.Command, args []string) {
		device.ListDevices()
	},
}
