package main

import (
	"fmt"
)

func main() {
	// Hello world
	fmt.Println("Hello, world with go")

	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println(pi, pi2)

	// Declaration variables entire
	base := 12
	var altura int = 14
	fmt.Println(base * altura)

	//Zero values => valores por defecto
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Calcular el area de un cuadrado

	var ladoA int = 10
	var ladoB int = 10
	var area = ladoA * ladoB

	fmt.Println("Area del cuadrado: ", area)

	// Operaciones aritemeticas sumas, restas, multiplicaciones y divisiones
	x := 50
	y := 100

	suma := x + y
	fmt.Println("Suma: ", suma)

	resta := x - y
	fmt.Println("Resta: ", resta)

	divisiones := x / y
	fmt.Println("Division: ", divisiones)

	multiplicaciones := x * y
	fmt.Println("Multiplicaciones: ", multiplicaciones)

	var count int = 10
	fmt.Println("Count: ", count)

	const myConst = "myConst"
	fmt.Println("myConst: ", myConst)

	// Controles de flujo

	if count > 5 {
		fmt.Println("Count is greater than 5")
	} else {
		fmt.Println("Count is less than or equal to 5")
	}

	// Arrays

	var myArray [3]int

	// bucles
	for i := 0; i < len(myArray); i++ {
	}

	/*  for _, value := range myArray {
	        fmt.Println(value)
	    }

	    // Structs

	    type Person struct {
	        Name string
	        Age  int
	    }

	    var person Person
	    person.Name = "Anderson"
	    person.Age = 30

	    fmt.Println("Person Name: ", person.Name)
	    fmt.Println("Person Age: ", person.Age)

	    // Funciones

	    func greet(name string) string {
	        return "Hello, " + name
	    }

	    fmt.Println(greet("Anderson"))

	    // Interfaces

	    type Animal interface {
	        Sound() string
	    }

	    type Dog struct{} */

}
