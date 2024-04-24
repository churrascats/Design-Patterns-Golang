package hospital

import "fmt"

type Cashier struct {
	next Handler
}

func (c *Cashier) SetNext(next Handler) {
	c.next = next
}

func (c *Cashier) Execute(p *Patient) {
	if p.PaymentDone {
		fmt.Println("Patient already paid the bill", p.Name)
		return
	} else {
		fmt.Println("Patient paying the bill", p.Name)
		p.PaymentDone = true
	}

	fmt.Println("Treatment done.")

}
