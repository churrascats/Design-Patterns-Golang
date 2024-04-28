package invoker

import "command/command"

type Button struct {
	SwitchOnOffCommand command.Command
}

func (b *Button) PressOnOff() {
	b.SwitchOnOffCommand.Execute()
}

func (b *Button) SetCommand(c command.Command) {
	b.SwitchOnOffCommand = c
}
