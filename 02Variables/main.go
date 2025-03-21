package main

import "fmt"

/*

Key Points to Remember
1. Use := only inside functions
2. Use var for global declaration
3. const can be used globally with var
4. In Go, every declared variable must be used, or the compiler will throw an error


Variable Scopes
1. Global variables : declared outside the function ; accessible throughout the package
2. Local variables  : declared inside the function ; accessible within that function
3. Block variables  : declared inside a block (like if , for) ; limited for that block

Default values
-int ->0  float -> 0.0  bool ->false  string -> ""(empty string)
-Pointers maps slices channels -> nil

Constants
-value must be assigned at the time of declaration and value cannot be changed later
- They can only hold primitive data types like int float string bool

*/

const Pi = 3.14
const Greeting = "Hello"

var globalVar = "I am global"

//fmt.Println("global var:",globalVar)

// message2 := "Hello, World!"  -->compilation error

func main() {
	var name string = "John Cena"
	var age int = 30

	country := "India" //shorthand

	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("country:", country)

	fmt.Printf("Type of name: %T\n", name) // Output: Type of name: string
	fmt.Printf("Type of age: %T\n", age)   // Output: Type of age: int
	fmt.Printf("Type of country: %T\n", country)

	fmt.Println("pi:", Pi)
	fmt.Println("greeting:", Greeting)

	fmt.Println("global var", globalVar)

	var i int
	var f float64
	var b bool
	var s string

	fmt.Println("int:", i)
	fmt.Println("float:", f)
	fmt.Println("bool:", b)
	fmt.Println("string", s)

	// var a, be, c int = 1, 2, 3
	// var x, y = "hello", true

	// d, e, f := 1, 2, 3
	// g, h := "hello", true

}
