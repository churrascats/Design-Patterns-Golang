package main

import "chain_of_responsability/hospital"

func main() {

	patient := hospital.Patient{
		Name: "Robertinho",
	}

	reception := &hospital.Reception{}
	doctor := &hospital.Doctor{}
	medical := &hospital.Medical{}
	cashier := &hospital.Cashier{}

	reception.SetNext(doctor)
	doctor.SetNext(medical)
	medical.SetNext(cashier)

	reception.Execute(&patient)
}
