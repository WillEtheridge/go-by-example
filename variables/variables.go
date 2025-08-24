/*
* 3. Variables
*
* In Go, variables are explicitly declared
* and used by the compiler to e.g
* check type-correctness of function calls
 */

package main

import "fmt"

// var declares 1 or more variable
var a = "initial"

// You can declare multiple variables at once
var b, c int = 1, 2

// Go will infer the type of initialized variables
var d = true

// Variables declared without a corresponding initialization are zero valued
// For example the zero value for an int is 0
var e int

func main() {
	fmt.Println(a)    // initial
	fmt.Println(b, c) // 1 2
	fmt.Println(d)    // true
	fmt.Println(e)    // 0

	// The := syntax is shorthand for declaring and initializating a variable
	// This syntax is only available inside functions
	f := "apple"
	fmt.Println(f) // apple
}
