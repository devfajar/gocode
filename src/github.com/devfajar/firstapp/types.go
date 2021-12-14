package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Bool Types
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c", c)
	d := a || b
	fmt.Println("d", d)

	// Integer Types
	// there is Int32 and Int64
	// use generally Int because its represent Integer depending on how the program perform
	var t int = 89
	f := 95
	fmt.Println("value of t is", t, "and f is", f)
	fmt.Printf("type of  t is %T, size of t is %d", t, unsafe.Sizeof(t))
	fmt.Printf("\ntype of f is %T, size of f is %d", f, unsafe.Sizeof(f))


	// Float Types
	g, h := 5.67, 8.97
	fmt.Printf("type of g %T h %T\n", g, h)
	sum := g + h
	diff := g - h
	fmt.Println("sum", sum, "diff", diff)

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)

	// Complex Types
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum: ", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)

	// String Types
	first := "Fajar Nur"
	last := "Trengginas"
	name := first +  " " +last
	fmt.Println("My name is", name)

	// Types convesion
	i := 55 		// int
	j := 67.8 		// float64
	added := i + int(j)	// int + float64 not allowed
	fmt.Println(added)


	k := 50
	var p float64 = float64(k) // this statement
	fmt.Println("p",p)


}
