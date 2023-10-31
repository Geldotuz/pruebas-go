package concesionario

import (
	"errors"
	"fmt"
	"prueba/carro"
)

type Concessionaire struct {
	Name    string
	Brand   string
	Capital float64
	Address string
	Cars    []carro.Auto
}

func NewConcessionaire(name string, brand string, capital float64, address string) *Concessionaire {
	return &Concessionaire{
		Name:    name,
		Brand:   brand,
		Capital: capital,
		Address: address,
	}
}

func (c *Concessionaire) SoldCar(i int) {
	c.Cars[i].Sold = true
}

func (c *Concessionaire) Register(a carro.Auto) {
	c.Cars = append(c.Cars, a)
}

func (c *Concessionaire) Charge(price float64) {
	//c.Capital + =a.price
	c.Capital += price
}

func (c *Concessionaire) Sell(price float64, color string) error {
	var auto carro.Auto
	var indexCar int
	for i, a := range c.Cars {
		if a.Sold {
			continue
		}

		if price == a.Price {
			auto = a
			indexCar = i
			break
		}
	}

	//preguntar y error

	if auto.Name == "" {
		return errors.New("no se encontro ningun carro a ese precio")
	}

	if auto.Color == color {
		//vendelo
		fmt.Print(auto)
		c.SoldCar(indexCar)
		c.Charge(auto.Price)
	} else if auto.Color != color {
		//pintalo
		//vendes
		fmt.Printf("Se pinto el auto de color %s a %s \n", auto.Color, color)
		auto.Colored(color)
		auto.Price += 200
		fmt.Printf("El precio del %s aumento en 200 y ahora cuesta %f \n", auto.Name, auto.Price)
		fmt.Println(auto)
		c.SoldCar(indexCar)
		c.Charge(auto.Price)

	}
	fmt.Printf("Capital de concesionario %f \n", c.Capital)
	return nil
}
