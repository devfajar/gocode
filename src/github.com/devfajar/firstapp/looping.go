package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++{
		if i > 5 {
			break // loop is terminated if i > 5
		}
		fmt.Printf("%d", i)
	}
	fmt.Printf("\n line after for loop \n")


	// loop with cotinue
	for j := 1; j <= 10; j++{
		if j%2 == 0 {
			continue
		}
		fmt.Println("%d", j)
	}

	// nested for loops
	n := 5
	for l := 0; l < n; l++{
		for k := 0; k <= l; k++{
			fmt.Printf("*")
		}
		fmt.Println()
	}

	// looping labels
	for a := 0; a < 3; a++ {
		for b := a; b < 4; b++{
			fmt.Printf("a = %d, j = %d\n", a, b)
		}
	}
}
