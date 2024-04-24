package hospital

import "fmt"

type Medical struct {
	next Handler
}

func (m *Medical) SetNext(next Handler) {
	m.next = next
}

func (m *Medical) Execute(p *Patient) {
	if p.MedicineDone {
		fmt.Println("Patient already medicated", p.Name)
		return
	} else {
		fmt.Println("Medicating the patient", p.Name)
		p.MedicineDone = true
	}

	m.next.Execute(p)

}
