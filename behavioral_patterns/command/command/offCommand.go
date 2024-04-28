package command

import (
	"command/devices"
)

type OffCommand struct {
	Device devices.Device
}

func (c *OffCommand) Execute() {
	c.Device.Off()
	c.Device.ShowStatus()
}
