package main

import (
	"fmt"
)

// slice function
func subtactOne(numbers []int){
	for i := range numbers{
		numbers[i] -= 2
	}
}

func countries() []string{
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) // copied neededCountries to countriesCpy
	return countriesCpy
}

func main() {
	a := [5]int{76, 75, 88, 87, 89}
	var b []int = a[1:4] // creates a slice for a[1] to a[3]
	fmt.Println(b)


	// another example
	c := []int{6, 7, 8} // creates an array and returns a slice references
	fmt.Println(c)

	// modifying a slice
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before ", darr)
	for i := range dslice{
		dslice[i]++
	}
	fmt.Println("array after", darr)

	// another slice
	numa := [3]int{78, 79, 80}
	nums1 := numa[:] // create a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", nums1)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", nums2)

	// fruits array
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "watermelon", "pineapple", "chikooo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice))


	// create slice using make
	i := make([]int, 5, 5)
	fmt.Println(i)

	// appending to slice
	cars := []string{"Ferrari", "Honda", "Ford",}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))

	// example
	var names []string // zero value of a slice
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents: ", names)
	}

	// veggies
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food : ", food)

	// slicing with function
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)
	fmt.Println("slice after function call", nos)


	// multidimensional slices
	pls := [][]string {
            {"C", "C++"},
            {"JavaScript"},
            {"Go", "Rust"},
            }
    for _, v1 := range pls {
        for _, v2 := range v1 {
            fmt.Printf("%s ", v2)
        }
        fmt.Printf("\n")
    }

	countriesNeeded := countries()
	fmt.Println(countriesNeeded)


}
