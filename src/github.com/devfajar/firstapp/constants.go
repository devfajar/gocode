package main

import (
	"fmt"
	//"math"
)

func main() {
	const (
		name = "john"
		age = 50
		country = "canada"
	)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(country)

	// cannot assigned a contrant
	//const b = 90
	//b = 88
	//fmt.Println(b)

	// example with math
	// var c = math.Sqrt(4) // allowed
	// const d = math.Sqrt(4) // not allowed

	// try
	const n = "sam"
	var named = n
	fmt.Printf("type %T value %v", named, named)

	// Go is a strongly typed language. Mixing types during the assignment is not allowed.
	//var defaultName = "Sam" // allowed
	//type myString string
	//var customName myString = "Sam" // allowed
	// customName = defaultName // not allowed
	//

	// boolean constant
	//const trueConst = true
	//type myBool bool
	//var defaultBool = trueConst // allowed
	//var customBool myBool = trueConst // allowed
	// defaultBool = customBool // not allowed

	// Numeric constants
	const e = 60
	var intVar int = e
	var int32Var int32 = e
	var float64Var float64 = e
	var complex64Var complex64 = e
	fmt.Println("intVar", intVar, "int32Var", int32Var, "float64Var", float64Var, "complex64Var", complex64Var)

	// another example numeric constants
	var i = 5
	var f = 5.6
	var m = 5 + 6i
	fmt.Println("i's type is %T, f's type is %T, m type is %T", i, f, m)


	// Numeric expressions
	var x = 5.9 / 8
	fmt.Printf("x's type is %T and value is %v", x, x)
}
