package main

import (
	"fmt"
	"math"
)

func main(){
	var age int //variable declaration
	fmt.Println("My age is ", age)
	age = 21
	fmt.Println("My age is ", age)
	age = 66
	fmt.Println("My age is ", age)

	// with initial value
	var num int = 44
	fmt.Println("num is", num)

	// multiple variable declaration
	var width, height int = 1024, 768
	fmt.Println("width is ", width, "height is", height)

	// shorthand variable declaration
	count := 10
	fmt.Println("count is", count)

	// multiple shorthand variable
	name, age := "fajar", 21
	fmt.Println("myname is", name, "age is", age)


	//try
	a, b := 20, 30 // declares variable a and b
	fmt.Println("a is", a, "b is", b)
	b, c :=  40, 50 // b is declared but c is already new
	fmt.Println("b is", b, "c is", c)
	b, c = 80, 90 // assign new values to already declared variable b and c
	fmt.Println("change b is", b, "c is", c)


	// but if we declare like this
	//  a, b := 50, 40
	//  fmt.Println("a is", a, "b is", b)
	//	a, b := 60, 30
	// there would be an error thats no new variable on the left side :=
	f, g := 145.8, 543.8
	j := math.Min(f, g)
	fmt.Println("Minimum value is", j)


}
