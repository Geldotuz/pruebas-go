package main

import (
	"fmt"
	"prueba2/parque"
)

func main() {
	park := parque.NewPark(3, "Quito")
	err := park.Enter(parque.Person{Name: "John", Age: 30})
	if err != nil {
		fmt.Println(err)
	}
	err = park.Enter(parque.Person{Name: "calo", Age: 12})
	if err != nil {
		fmt.Println(err)
	}
	err = park.Enter(parque.Person{Name: "cris", Age: 25})
	if err != nil {
		fmt.Println(err)
	}
	park.ReleaseCar()
	err = park.Enter(parque.Person{Name: "angel", Age: 100})
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(park)
	park.Report()
}
