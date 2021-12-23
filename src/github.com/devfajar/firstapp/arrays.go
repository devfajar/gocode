package main

import (
	"fmt"
)

func changeLocal(num [5]int){
	num[0] = 55
	fmt.Println("inside fucntion ", num)
}

func printarray(a [3][2]string){
	for _, v1 := range a{
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func main() {
	var a [3]int // int array with length 3
	a[0] = 12 // array index start at 0
	a[1] = 78
	a[2] = 50
	fmt.Println(a)

	// short version making array
	b := [3]int{12, 78, 50}
	fmt.Println(b)

	c := [3]int{12}
	fmt.Println(c)


	// ignoring length of array
	d := [...]int{14, 78, 99}
	fmt.Println(d)

	// cant be resizing
	// e := [3]int{5, 99, 90}
	// var f [5] int
	// f = e //not possible sice [3]int and 5[int]

	// another example arrays
	g := [...]string{"USA", "China", "India", "Germany", "France"}
	h := g // a copy of g assigned to h
	h[0] = "Singapore"
	fmt.Println("g is ", g)
	fmt.Println("h is ", h)

	// working with function changeLocal
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) // num is passed by value
	fmt.Println("after passing to function ", num)

	// length of array
	l := [...]float64{64.7, 89.8, 21, 78}
	fmt.Println("length of l is ", len(l))

	// iterating arrays using range
	k := [...]float64{67.7, 90.0, 22, 77}
	for i := 0; i < len(k); i++ {
		fmt.Printf("%d th element of k is %.2f\n", i, k[i])
	}

	// better iterate
	j := [...]float64{66.8, 70.9, 22, 66.6}
	sum := float64(0)
	for i, v := range j { // range returns both the index and value
		fmt.Printf("%d the element of j is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("sum of all elements of j ",sum)


	// Multipledimensional array
	m := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"},
	}
	printarray(m)
	var n [3][2]string
	n[0][0] = "apple"
	n[0][1] = "samsung"
	n[1][0] = "microsoft"
	n[1][1] = "google"
	n[2][0] = "AT&T"
	n[2][1] = "T-mobile"
	fmt.Printf("\n")
	printarray(n)
}
