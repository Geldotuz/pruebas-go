package main

import (
	"fmt"
	"prueba3/players"
)

func main() {
	angel := players.NewPlayer("angel", "martin", 30)

	fmt.Printf("%v %v tiene %d de edad ", angel.Name, angel.LastName, angel.Age)

	angel.Cumpleaños()

	angel.CambioApellido("ruiz")

	fmt.Printf("%v %v tiene %d de edad ", angel.Name, angel.LastName, angel.Age)
}
