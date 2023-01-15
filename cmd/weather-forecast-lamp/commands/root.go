package commands

import (
	"fmt"
	"os"

	"github.com/kironono/weather-forecast-lamp/internal/device"
	"github.com/kironono/weather-forecast-lamp/internal/weather"
	"github.com/nasa9084/go-switchbot"
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

	rootCmd.PersistentFlags().StringP("switchbot-open-token", "", "", "Your SwitchBot open token")
	rootCmd.PersistentFlags().StringP("switchbot-secret-key", "", "", "Your SwitchBot secret key")
	rootCmd.PersistentFlags().StringP("switchbot-device-id", "", "", "Target control device ID")
	rootCmd.PersistentFlags().StringP("openweather-app-id", "", "", "Your OpenWeatherMap app ID")
	rootCmd.PersistentFlags().StringP("openweather-city", "", "", "OpenWeatherMap weather forecast city")

	viper.BindPFlag("switchbot_open_token", rootCmd.PersistentFlags().Lookup("switchbot-open-token"))
	viper.BindPFlag("switchbot_secret_key", rootCmd.PersistentFlags().Lookup("switchbot-secret-key"))
	viper.BindPFlag("switchbot_device_id", rootCmd.PersistentFlags().Lookup("switchbot-device-id"))
	viper.BindPFlag("openweather_app_id", rootCmd.PersistentFlags().Lookup("openweather-app-id"))
	viper.BindPFlag("openweather_city", rootCmd.PersistentFlags().Lookup("openweather-city"))
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
	rootCmd.RunE = execUpdateCommandF
	return rootCmd.Execute()
}

func execUpdateCommandF(command *cobra.Command, args []string) error {
	token := viper.GetString("switchbot_open_token")
	secret := viper.GetString("switchbot_secret_key")
	deviceId := viper.GetString("switchbot_device_id")
	openWeatherAppId := viper.GetString("openweather_app_id")
	city := viper.GetString("openweather_city")

	sc := switchbot.New(token, secret)

	cb := device.NewColorBulb(sc, deviceId)
	pop, err := weather.GetProbabilityOfPrecipitation(openWeatherAppId, city)
	if err != nil {
		return fmt.Errorf("get pop: %w", err)
	}
	return cb.Update(pop)
}
