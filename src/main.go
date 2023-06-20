package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	pk "GoBasic/src/mypackage"
)

func main() {
	var carro pk.CarPublic
	carro.Brand = "Buggati"
	carro.Year = 2020
	fmt.Println(carro)

	pk.PrintMessage()
}

type car struct {
	brand string
	year  int
}

func structs() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	//Otra manera
	var otherCar car
	otherCar.brand = "Koenigsegg"
	fmt.Println(otherCar)
}

func maps() {
	m := make(map[string]int)
	m["Jose"] = 14
	m["Maria"] = 56

	fmt.Println(m)

	//Recorrido de map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar un valor
	value, ok := m["Josep"]
	fmt.Println(value, ok)

	value2, ok2 := m["Jose"]
	fmt.Println(value2, ok2)
}

func isPalindromo(text string) {
	var textReverse string
	var textMinus string = strings.ToLower(text)

	for i := len(textMinus) - 1; i >= 0; i-- {
		textReverse += string(textMinus[i])
	}

	if textMinus == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func sliceRange() {
	slice := []string{"hola", "que", "hace"}

	for _, valor := range slice {
		fmt.Println(valor) // lo mismo para el indice
	}

}

func arraysAndSlice() {
	//Array (inmutables)
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	//Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(slice, len(slice), cap(slice))

	//metodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4]) //[a es inclusivo: b es exclusivo]
	fmt.Println(slice[4:])

	//Append
	slice = append(slice, 9)
	fmt.Println(slice)

	//Append para listas
	newSlice := []int{10, 11, 12}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}

func keyWords() {
	//Defer
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		//Break
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}
}

func switchConditional() {
	switch modulo := 45 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	//Sin condicion

	value := 50
	switch {
	case value > 100:
		fmt.Println("Es mayor que 100")
	case value < 0:
		fmt.Println("Es menor que 0")
	default:
		fmt.Println("No hay condicion")
	}
}

func parDetection(a int) bool {
	var detector bool

	if a%2 == 0 {
		detector = true
		fmt.Println("Es par")
		return detector
	} else {
		detector = false
		fmt.Println("No es par")
		return detector
	}
}

func stringToInt() {
	value, err := strconv.Atoi("53")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value:", value)
}

func areaCuadrado(a, b int) int {
	return a * b
}

func areaTrapecio(b, B, h int) int {
	return (B - b) * h / 2
}

func areaCirculo(r int) int {
	pi := 3.141592
	return r * int(pi*pi)
}

func ifConditional() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	//With and
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es igual y no es cero")
	}

	//With or
	if valor1 == 1 || valor2 == 0 {
		fmt.Println("Es uno o cero")
	}
}

func forForever() {
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}
}

func forWhile() {
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}
}

func forConditional() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func returnValue(a int) int {
	return a * 2
}

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func manejoVariables() {
	//Declaracion de constantes
	const pi float64 = 3.141592
	const pi2 = 3.14

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	//Declaracion de Variables
	base := 12          //Declaracion sin tipo de dato
	var altura int = 14 //Declaracion entera (Tipo de variable + valor)
	var area int        //Declaracion sin asignacion

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	//Area Cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area Cuadrardo:", areaCuadrado)

	x := 10
	y := 50

	//Suma
	result := x + y

	fmt.Println("Suma:", result)

	//Resta
	result = y - x
	fmt.Println("Resta:", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multiplicación:", result)

	//Division
	result = y / x
	fmt.Println("Division:", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	//Incremental
	x++
	fmt.Println("Incremental:", x)

	//Decremental
	x--
	fmt.Println("Decremental:", x)
}

func fmtFunction() {
	//PAQUETE FMT

	//Declaracion de variables

	helloMessage := "Hello"
	worldMessage := "World"

	//Println
	fmt.Println(helloMessage, worldMessage)

	//Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos) // %s = Strings, %d = numeros
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos) //%v = cualquier tipo de variable

	//Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	//Tipo de datos con fmt
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T", cursos)

}
