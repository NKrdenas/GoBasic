package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo"
)

func main() {
	// go modules
	e := echo.New()
	//Ruta
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

//Range, Close y Select

func rangeCloseSelect() {
	c := make(chan string, 2)
	c <- "Mensaje 1"
	c <- "Mensaje 2"

	fmt.Println(len(c), cap(c))

	//Range y Close
	close(c)

	// c <- "Mensaje 3"

	for message := range c {
		fmt.Println(message)
	}

	// Select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("mensaje1", email1)
	go message("mensaje2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}

}

func message(text string, c chan string) {
	c <- text
}

//Channels

func channels() {
	c := make(chan string, 1)

	fmt.Println("Hello")

	go say2("Bye", c)

	fmt.Println(<-c)
}

func say2(text string, c chan<- string) {
	c <- text
}

//Go routines

func goRoutine() {
	var wg sync.WaitGroup

	fmt.Println("Hello")
	wg.Add(1)

	go say("world", &wg)

	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("Adios")

	time.Sleep(time.Second * 1)
}

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(text)
}

//Interfaces

func interfaces() {
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 3, altura: 5}

	calculate(myCuadrado)
	calculate(myRectangulo)

	//Lista de interfaces
	myIntercace := []interface{}{"Hola", 12, 4.9}
	fmt.Println(myIntercace...)
}

type figuras2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func calculate(f figuras2D) {
	fmt.Println("Area:", f.area())
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

//stringers

func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB de Disco, y es una %s", myPC.ram, myPC.disk, myPC.brand)
}

func stringers() {
	myPC := pc{ram: 16, brand: "msi", disk: 100}
	fmt.Println(myPC)
}

type pc struct {
	ram   int
	disk  int
	brand string
}

//pointers

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRam() {
	myPC.ram = myPC.ram * 2
}

func pointers() {
	a := 50
	b := &a
	fmt.Println(a)
	fmt.Println(b)  // Direccion de memoria de a
	fmt.Println(*b) // Valor de la direccion de memoria apuntada

	*b = 100
	fmt.Println(a) // *b = a, apuntan al mismo lugar de la memoria
	myPC := pc{ram: 16, disk: 20, brand: "msi"}
	fmt.Println(myPC)
	myPC.ping()

	fmt.Println(myPC)
	myPC.duplicateRam()

	fmt.Println(myPC)
	myPC.duplicateRam()
	fmt.Println(myPC)
}

// func packageStructs() {
// 	var carro pk.CarPublic
// 	carro.Brand = "Buggati"
// 	carro.Year = 2020
// 	fmt.Println(carro)

// 	pk.PrintMessage("Hola Platzi")

// 	// pk.printMessage("Hola") No se puede utilizar ya que es privado
// }

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
