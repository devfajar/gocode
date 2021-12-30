package main


import (
	"fmt"
	"unicode/utf8"
)

func printBytes(s string){
	fmt.Printf("Bytes :")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
		}
}
func printChars(s string){
	fmt.Printf("Characters :")
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
		}
}

func charsAndBytePosition(s string){
	for index, rune := range s{
		fmt.Printf("%c starts at bye %d\n", rune, index)
	}
}

// String Comparison Function
func compareStrings(str1 string, str2 string){
	if str1 == str2 {
		fmt.Printf("%s and %s are equal\n", str1, str2)
		return
	}
	fmt.Printf("%s and %s are not equal\n", str1, str2)
}

// Mutate Function
func mutate(s []rune) string {
	s[0] = 'a' //any valid unicode character within single quote is a rune
	return string(s)
}

func main() {
	name := "Hello World"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)
	fmt.Printf("\n")
	name = "Señor"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)
	fmt.Printf("\n")
	charsAndBytePosition(name)

	// Create String from a slice of byte
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	byteSlice = []byte{67, 97, 102, 195, 169}//decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	str := string(byteSlice)
	fmt.Println(str)

	// Create String from a slice of rune
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str = string(runeSlice)
	fmt.Println(str)

	// String length
	/* The RuneCountInString(s string) (n int) function of the utf8 package
	 can be used to find the length of the string. */
	word1 := "Señor"
	fmt.Printf("String: %s\n", word1)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word1))
	fmt.Printf("Number of bytes: %d\n", len(word1))

	fmt.Printf("\n")
	word2 := "Pets"
	fmt.Printf("String: %s\n", word2)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word2))
	fmt.Printf("Number of bytes: %d\n", len(word2))

	// String Comparison
	string1 := "Go"
	string2 := "Go"
	compareStrings(string1, string2)

	string3 := "hello"
	string4 := "world"
	compareStrings(string3, string4)

	// String Concatenation
	string5 := "Go"
	string6 := "is awesome"
	result := string5 + " " + string6
	fmt.Println(result)

	// String Concatenation with sprintf
	string7 := "Go"
	string8 := "is awesome"
	result = fmt.Sprintf("%s %s", string7, string8)
	fmt.Println(result)

	// String is immutable
	// Strings are immutable in Go. Once a string is created it's not possible to change it.
	// h := "hello"
	// fmt.Println(mutate(h))
	h := "hello"
	fmt.Println(mutate([]rune(h)))
}
