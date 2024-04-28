package main

import (
	"command/command"
	"command/devices"
	"command/invoker"
)

func main() {

	button := invoker.Button{}
	var device devices.Device = &devices.Television{}

	turnOnCommand := command.OnCommand{Device: device}
	button.SetCommand(&turnOnCommand)
	button.PressOnOff()

	turnOffCommand := command.OffCommand{Device: device}
	button.SetCommand(&turnOffCommand)
	button.PressOnOff()
}
