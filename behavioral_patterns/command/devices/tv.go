package devices

import "fmt"

type Television struct {
	IsON bool
}

func (t *Television) On() {
	t.IsON = true
}

func (t *Television) Off() {
	t.IsON = false
}

func (t *Television) ShowStatus() {
	fmt.Println("Tv turned ", t.IsON)
}
