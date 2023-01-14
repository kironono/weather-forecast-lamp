package device

import (
	"context"

	"github.com/fatih/color"
	"github.com/nasa9084/go-switchbot"
	"github.com/rodaine/table"
	"github.com/spf13/viper"
)

func ListDevices() {
	token := viper.GetString("switchbot_open_token")
	secret := viper.GetString("switchbot_secret_key")

	c := switchbot.New(token, secret)
	pdev, _, _ := c.Device().List(context.Background())

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("ID", "Type", "Name")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, d := range pdev {
		tbl.AddRow(d.ID, d.Type, d.Name)
	}

	tbl.Print()
}
