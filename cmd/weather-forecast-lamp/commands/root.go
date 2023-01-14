package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "weather-forecast-lamp",
		Short: "Weather forecast lamp",
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("switchbot-open-token", "", "", "Your SwitchBot open token.")
	rootCmd.PersistentFlags().StringP("switchbot-secret-key", "", "", "Your SwitchBot secret key.")

	viper.BindPFlag("switchbot_open_token", rootCmd.PersistentFlags().Lookup("switchbot-open-token"))
	viper.BindPFlag("switchbot_secret_key", rootCmd.PersistentFlags().Lookup("switchbot-secret-key"))
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory.
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".weather-forecast-lamp")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func Run(args []string) error {
	rootCmd.SetArgs(args)
	return rootCmd.Execute()
}
