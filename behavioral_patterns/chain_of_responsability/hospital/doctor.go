package hospital

import "fmt"

type Doctor struct {
	next Handler
}

func (d *Doctor) SetNext(next Handler) {
	d.next = next
}

func (d Doctor) Execute(p *Patient) {
	if p.DoctorCheckUpDone {
		fmt.Println("Doctor checkup already done", p.Name)
		return
	} else {
		fmt.Println("Doctor checking up the patient", p.Name)
		p.DoctorCheckUpDone = true
	}

	d.next.Execute(p)

}
