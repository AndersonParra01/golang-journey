package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world with go")

	// Variables

	var name string = "hello"
	fmt.Println(name + " " + "Anderson")

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
