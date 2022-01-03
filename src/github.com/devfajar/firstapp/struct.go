package main

import (
	"fmt"
)

// Declaring struct
type employee struct{
	firstName string
	lastName string
	age int
	salary int
	country string
}

func main(){
	emp1 := employee{
		firstName: "SAM",
		age: 25,
		lastName: "Anderson",
		salary: 12000,
		country: "USA",
	}
	// creating struct without specifying field names
	emp2 := employee{"Thomas", "Paul", 29, 800, "Brazil"}

	// Create anonymous struct
	emp3 := struct {
		firstName string
		lastName string
		age int
		salary int
	}{
		firstName: "Andreah",
		lastName: "Nikola",
		age: 31,
		salary: 3000,
	}

	// Accessing individuals field of a struct
	emp4 := employee{
		firstName: "Fajar",
		lastName: "Test",
		age: 21,
		salary: 6000,
	}

	// Zero Value of a struct
	var emp5 employee // zero valued struct

	// Pointer to struct
	emp6 := &employee {
		firstName: "John",
		lastName: "Cornor",
		age: 55,
		salary: 5000,
	}

	// Anonymous Field
	type Address struct {
		city string
		state string
	}

	type Person struct {
		name string
		age int
		address Address
	}

	// Nested struct

	fmt.Println("Employee 1 ", emp1)
	fmt.Println("Employee 2 ", emp2)
	fmt.Println("Employee 3 ", emp3)


	// print individuals field of struct emp4
	fmt.Println("Individual field of emp4")
	fmt.Println("First Name: ", emp4.firstName)
	fmt.Println("Last Name: ", emp4.lastName)
	fmt.Println("Age: ", emp4.age)
	fmt.Printf("Salary: $%d\n", emp4.salary)
	emp4.salary = 6500
	fmt.Printf("New Salary: $%d\n", emp4.salary)

	fmt.Printf("\n")

	fmt.Println("First Name: ", emp5.firstName)
	fmt.Println("Last Name: ", emp5.lastName)
	fmt.Println("Age: ", emp5.age)
	fmt.Printf("Salary: $%d\n", emp5.salary)

	fmt.Printf("\n")

	fmt.Println("First Name: ", (*emp6).firstName)
	fmt.Println("Age: ", (*emp6).age)

	// struct person
	fmt.Println("New Struct Person")
	p1 := Person{
		name: "naveen",
		age: 50,
	}


	fmt.Println(p1.name)
	fmt.Println(p1.age)

	// Nested Struct
	p := Person{
		name: "death",
		age: 60,
		address: Address{
			city: "Chicago",
			state: "Illnois",
		},
	}

	fmt.Println("Name : ", p.name)
	fmt.Println("Age : ", p.age)
	fmt.Println("City :", p.address.city)
	fmt.Println("State :", p.address.state)

}
