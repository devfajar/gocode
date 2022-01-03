package main

import (
	"fmt"
)
// pass pointer value to pointer
func change(val *int) {
	*val = 55
}

// returning pointer from a function
func hello() *int {
	i := 5
	return &i
}

// func modify(arr *[3]int){
//	arr[0] = 90
// }

func modify(sls []int){
	sls[0] = 90
}

func main() {
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is ", a)

	// Zero value pointer
	c := 25
	var d *int
	if d == nil {
		fmt.Println("d is ", d)
		d = &c
		fmt.Println("d after initialization is", d)
	}

	// Creating pointers using the new function
	size := new(int)
	fmt.Printf("Size value is %d, type is %T, and address is %v\n", size, size, size)
	*size = 85
	fmt.Println("New size value is ", *size)

	// Dereferencing a pointer
	e := 255
	f := &e
	fmt.Println("address of f is ", f)
	fmt.Println("value of e is ", *f)
	*f++
	fmt.Println("new value of e is ", e)


	// Passing pointer to a function
	g := 58
	fmt.Println("value of g before function call is ", g)
	h := &g
	change(h)
	fmt.Println("value of g after function call is ", g)

	// returning pointer from a fnction
	j := hello()
	fmt.Println("Value of j ", *j)

	// Do not pass a pointer to an array as an argument to a function. Use slice instead.
	// k := [3]int{89, 90, 91}
	// modify(&k)
	// fmt.Println(k)

	// with slices
	k := [3]int{89, 90, 91}
	modify(k[:])
	fmt.Println(k)

	// Go does not support pointer arithmetic
	q := [...]int{109, 110, 111}
	p := &q
	p++
}
