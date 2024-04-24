package hospital

import "fmt"

type Reception struct {
	next Handler
}

func (r *Reception) SetNext(next Handler) {
	r.next = next
}

func (r Reception) Execute(p *Patient) {
	if p.RegistrationDone {
		fmt.Println("Reception already registered the patient ", p.Name)
		return
	} else {
		fmt.Println("Reception registering the patient", p.Name)
		p.RegistrationDone = true
	}

	r.next.Execute(p)

}
