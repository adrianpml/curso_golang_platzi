package main

import (
	"fmt"
	"log"
	"strconv"
)

func funcionNormal(message string) {
	fmt.Println(message)
}

func returnValue(a, b int) int {
	return a * b
}

func doubleReturnFunc(a int) (c, d int) {
	return a, a * 2
}

func main() {
	value3, value4 := doubleReturnFunc(2)
	fmt.Println("valores: ", value3, value4)
	fmt.Println(returnValue(1, 3))
	funcionNormal("Hola Mundo")
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

	// Rectángulo
	baseRectangulo := 20
	alturaRectangulo := 10

	areaRectangulo := baseRectangulo * alturaRectangulo

	fmt.Println("El Area del Rectángulo es :", areaRectangulo)

	// Circulo : AreaCirculo = pi por radio al cudrado
	const PI float64 = 3.14 // Constant
	var radioCirculo float64 = 10

	areaCirculo := PI * radioCirculo * radioCirculo

	fmt.Println("El Area del Circulo es :", areaCirculo)

	// Trapecio
	var baseUno float64 = 6
	var baseDos float64 = 15
	var alturaTrapecio float64 = 25

	areaTrapecio := ((baseUno + baseDos) * alturaTrapecio) / 2

	fmt.Println("El Area del Trapecio es :", areaTrapecio)

	helloMessage := "Hello"
	worldMessage := "world"

	// Println: Salto de Linea Automatico
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	// Con valores seguros
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	// Con valores inseguros
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%v tiene más de %v cursos\n", nombre, cursos)
	fmt.Println(message)

	// Tipo de datos:
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

	//If
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("es 2")
	}

	//With and
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	}

	//With or
	if valor1 == 0 || valor2 == 1 {
		fmt.Println("Es verdad")
	}

	//convertir texto a numero
	value, error := strconv.Atoi("53")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(value)

	//MAPS
	m := make(map[string]int)

	m["Jose"] = 14
	m["Juan"] = 15

	fmt.Println(m)

	//Recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar valor
	value, ok := m["Juan"]
	fmt.Println(value, ok)

}
