/*
* 2. Values
* Go has various value types
* including strings, integers, floats, booleans, etc.
 */

package main

import "fmt"

func main() {
	// Strings can be added with a +
	fmt.Println("go" + "lang") // golang

	// Integers
	fmt.Println("1+1 = ", 1+1) // 1+1 = 2

	// Floats
	fmt.Println("7.0/3.0 = ", 7.0/3.0) // 7.0/3.0 = 2.3333333333333335

	// Booleans with boolean operators
	fmt.Println(true && false) // false
	fmt.Println(true || false) // true
	fmt.Println(!true)         // false
}
