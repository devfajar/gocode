package main

import (
	"fmt"
	"math/rand"
)

func main() {
	finger := 6
	fmt.Printf("Finger %d is ", finger)
	switch finger{
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default: // if we input incorrect number switch will execute default
 		fmt.Println("incorrect finger number")
	}


	// Multiple expression case
	letter := "i"
	fmt.Printf("Letter %s is ", letter)
	switch letter{
	case "a", "e", "i", "u", "o": // multiple expression case
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}

	// expressionless switch
	num := 75
	switch { // expression is omitted
    case num >= 0 && num <= 50:
		fmt.Println("%d is greater than 0 and less than 50", num)
	case num >= 51 && num <= 100:
		fmt.Println("%d is greater than 51 and less than 100", num)
	case num >= 101:
		fmt.Println("%d is greater than 100", num)
	}

	// break
	switch number := -5;{
    case number < 50:
		if number < 0 {
			break
		}
		fmt.Printf("%d is lesser than 50\n", number)
		fallthrough
	case number < 100:
		fmt.Printf("%d is lesser than 100\n", number)
		fallthrough
	case number < 200:
		fmt.Println("%d is lesser than 200\n", number)
	}

	// break the outer top
randloop:
	for {
		switch i := rand.Intn(100);{
		case i%2 == 0:
			fmt.Println("generate even number %d ", i)
			break randloop
		}
	}
}
