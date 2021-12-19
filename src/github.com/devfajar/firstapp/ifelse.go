package main

import (
	"fmt"
)

func main() {
	/*
	num := 11
	if num%2 == 0{ // checks if number can divideby 2
		fmt.Println("The number", num, "is even")
	}else{
		fmt.Println("The number ", num, "is odd")
	}
	*/
	// if elseif else
	number := 99
	if number <= 50 {
		fmt.Println(number, "is less than or equal 50")
	}else if number >= 51 && number <= 100{
		fmt.Println(number, "is between 51 and 100")
	}else{
		fmt.Println(number, "is greater than 100")
	}

	// if with assignment
	if nums := 10; nums % 2 == 0{ // checks if number is even
		fmt.Println(nums,"is even")
	}else{
		fmt.Println(nums,"is odd")
	}

	// go idiomatic
	num := 10
	if num%2 == 0{ // checks if number is even
		fmt.Println(num,"is even")
		return
	}
	fmt.Println(num,"is odd")

}
