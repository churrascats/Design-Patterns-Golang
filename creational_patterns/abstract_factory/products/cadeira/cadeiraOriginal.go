package cadeira

import "fmt"

type CadeiraOriginal struct {
}

func (c CadeiraOriginal) Sentar() {
	fmt.Println("Sentando na cadeira original")
}
