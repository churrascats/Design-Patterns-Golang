package mesa

import "fmt"

type MesaOriginal struct {
}

func (c MesaOriginal) ColocarObjeto() {
	fmt.Println("Sentando na mesa original")
}
