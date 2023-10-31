package parque

import (
	"errors"
	"fmt"
)

type Person struct { //paquete
	Name string
	Age  int
}

type Park struct {
	People     map[string]int
	Location   string
	CarsParked int
	Capacity   int
}

func NewPark(capacity int, location string) *Park { //paquete
	return &Park{
		Capacity:   capacity,
		Location:   location,
		CarsParked: 0,
		People:     make(map[string]int), // map vacion
	}
}

func (p *Park) AddCars() { // no paquete
	p.CarsParked++
}

func (p *Park) AddPeople(name string, age int) { // Ingresamos personas
	p.People[name] = age
}

func (p *Park) ReleaseCar() {
	p.CarsParked--
}

func (p *Park) Report() {
	for name, age := range p.People {
		fmt.Printf("El nombre es %s y su edad %d \n", name, age)
	}
	fmt.Println("el numero total de persona es ", len(p.People))
}

func (p *Park) Enter(person Person) error { // no paquete
	if p.CarsParked < p.Capacity {
		p.AddCars()
		p.AddPeople(person.Name, person.Age)
	} else {
		return errors.New("the park is full")
	}
	return nil
}
