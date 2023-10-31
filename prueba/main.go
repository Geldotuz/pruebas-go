package main

import (
	"fmt"
	"prueba/carro"
	"prueba/concesionario"
)

func main() {
	car1 := carro.NewAuto("Corola", "model1", 1000, "red", false)
	car2 := carro.NewAuto("Carril", "model2", 500, "blue", false)
	car3 := carro.NewAuto("Ferrari", "model3", 2500, "green", false)
	car4 := carro.NewAuto("Trekinnon", "model4", 5000, "yellow", false)

	concesionario := concesionario.NewConcessionaire("Maresa", "Toyota", 5000, "Manosca")

	concesionario.Register(*car1)
	concesionario.Register(*car2)
	concesionario.Register(*car3)
	concesionario.Register(*car4)

	err := concesionario.Sell(500, "red")
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	err = concesionario.Sell(500, "red")
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	err = concesionario.Sell(2500, "green")
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	//concesionario.Sell(1000, "blue")
}
