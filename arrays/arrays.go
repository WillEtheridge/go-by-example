/*
* 8. Arrays
*
* In Go, an array is a numbered sequence of elements of a specific length.
* In typical Go code, slices are much more common;
* arrays are useful in some special scenarios
 */

package main

import "fmt"

func main() {
	/*
	 * Here we create an array that will hold exactly 5 ints.
	 * The type of elements and length are both part of the arrays type.
	 * By default, an array is zero-valued, which for ints means 0
	 */
	var a [5]int
	fmt.Println("emp:", a)

	/*
	* We can set a value at an index using the array[index] = value syntax,
	*  and get a value with array[index]
	 */
	a[4] = 100
	fmt.Println("set", a)
	fmt.Println("get", a[4])

	// The built in len returns the length of an array
	fmt.Println("len:", len(a))

	// Use this syntax to declare and initialize an array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// You can also have the compiler count the number of elements for you with ...
	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", c)

	// If you specify the index with :, the elements inbetween will be zeroed
	e := [...]int{100, 3: 400, 500}
	fmt.Println("idx:", e) // [100, 0, 0, 400, 500]

	// Array types are one-dimensional but you can compose types to build multi-dimensional data structures
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // 2d:  [[0 1 2] [1 2 3]]

	// You can create and initialize multi-dimensional arrays at once too
	twoDtwo := [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d2: ", twoDtwo)
}
