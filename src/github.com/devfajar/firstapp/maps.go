package main

import (
	"fmt"
)

func main() {
	employeeSalary := map[string]int {
		"steve" : 12000,
		"jamie" : 15000,
		"mike" : 9000,
	}
	fmt.Println("map before deletion ", employeeSalary)
	delete(employeeSalary, "steve")
	fmt.Println("map after deletion", employeeSalary)


	// Not done
	/* maps equality
	map1 := map[string]int{
		"one":1,
		"two":2,
	}
	map2 := map1
	if map1 == map2{

	}
	*/
}
