package main

import "fmt"

func main() {
	//Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	//Declaracion de variables enteras

	//Area cuadrado
	const baseCuadrado = 10
	var areaCuadrado = baseCuadrado * baseCuadrado

	fmt.Println("Area del cuadrado: ", areaCuadrado)

	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("El resultado de la suma es: ", result)

	result = y - x
	fmt.Println("El resultado de la resta es: ", result)

}
