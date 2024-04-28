package command

import (
	"command/devices"
)

type OnCommand struct {
	Device devices.Device
}

func (c *OnCommand) Execute() {
	c.Device.On()
	c.Device.ShowStatus()
}
