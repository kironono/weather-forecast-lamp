package device

import (
	"context"
	"fmt"

	"github.com/nasa9084/go-switchbot"
)

type ColorBulb struct {
	Client   *switchbot.Client
	DeviceID string
}

func NewColorBulb(client *switchbot.Client, deviceId string) *ColorBulb {
	return &ColorBulb{
		client,
		deviceId,
	}
}

func (cb *ColorBulb) Update(pop float64) error {
	ctx := context.Background()
	err := cb.Client.Device().Command(
		ctx,
		cb.DeviceID,
		switchbot.DeviceCommandRequest{
			Command:     "setBrightness",
			Parameter:   "100",
			CommandType: "command",
		})
	if err != nil {
		return err
	}

	err = cb.Client.Device().Command(
		ctx,
		cb.DeviceID,
		switchbot.DeviceCommandRequest{
			Command:     "setColor",
			Parameter:   cb.colorLabel(pop),
			CommandType: "command",
		})
	if err != nil {
		return err
	}

	err = cb.Client.Device().Command(
		ctx,
		cb.DeviceID,
		switchbot.DeviceCommandRequest{
			Command:     "turnOn",
			CommandType: "command",
		})
	if err != nil {
		return err
	}

	return nil
}

func (cb *ColorBulb) colorLabel(p float64) string {
	r := 255
	g := 160 + int(p*95)
	b := 70 + int(p*95)
	return fmt.Sprintf("%d:%d:%d", r, g, b)
}
