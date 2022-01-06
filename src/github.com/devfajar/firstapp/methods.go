package main

import(
	"fmt"
	"math"
)

type Employee struct{
	name string
	salary int
	currency string
}

// new struct
type Rectangle struct {
	length int
	width int
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64{
	return math.Pi * c.radius * c.radius
}

func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

/*
Method with value receiver
*/
func (e Employee) changeName(newName string) {
    e.name = newName
}

/*
Method with pointer receiver
*/
func (e *Employee) changeAge(newAge int) {
    e.age = newAge
}

func main() {
	emp1 := Employee {
		name: "Sam Adolf",
		salary: 5000,
		currency: "$",
	}
	emp1.displaySalary() // calling display salary methods

	// Methods vs Function
	r := Rectangle {
		length : 10,
		width : 5,
		}
	fmt.Printf("Area of Rectangle %d\n", r.Area())

	c := Circle {
		radius : 12,
		}
	fmt.Printf("Area of Circle %f", c.Area())

	// Pointer Receivers vs Value Receivers

}
