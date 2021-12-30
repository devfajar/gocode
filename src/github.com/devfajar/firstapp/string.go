package main


import (
	"fmt"
)
	// Accessing individual bytes of a string
	// Since a string is a slice of bytes,
	// it's possible to access each byte of a string.

func printBytes(s string){
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

	// Accessing Individual Characters of a string
func printChars(s string){
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++{
		fmt.Printf("%c ", s[i])
	}
}

func main() {
	// simple string
	name := "Hello World!"
	fmt.Printf("String: %s\n", name)
    printChars(name)
    fmt.Printf("\n")
    printBytes(name)
    fmt.Printf("\n\n")
    name = "Señor"
    fmt.Printf("String: %s\n", name)
    printChars(name)
    fmt.Printf("\n")
    printBytes(name)
}
