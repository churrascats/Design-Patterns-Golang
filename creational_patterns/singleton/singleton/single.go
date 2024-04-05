package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}
var singleInstance *single

type single struct{}

func GetInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleInstance == nil {
			fmt.Println("Creating a new instance")
			singleInstance = &single{}
		} else {
			fmt.Println("Instance already created")
		}

	} else {
		fmt.Println("Instance already created")
	}

	return singleInstance
}
