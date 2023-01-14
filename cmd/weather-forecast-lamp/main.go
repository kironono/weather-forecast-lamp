package main

import (
	"os"

	"github.com/kironono/weather-forecast-lamp/cmd/weather-forecast-lamp/commands"
)

func main() {
	if err := commands.Run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
