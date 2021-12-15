package main


import "fmt"

// example function
func calculateBill(price, no int) int{
	var totalPrice = price * no
	return totalPrice
}
// multiple return values
func rectProps(length, width float64)(float64, float64){
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}
// named return values
func namedReturn(panjang, lebar float64)(daerah, perameter float64){
	daerah = panjang * lebar
	perameter = (panjang + lebar) * 2
	return daerah, perameter
}

func main() {
	// function calculate Bill
	price, no := 90, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("Total Price is", totalPrice)

	// function rectprops
	area, perimeter := rectProps(10.8, 5.6)
	fmt.Println("Area %f perimeter %f", area, perimeter)

	//named return function
	daerah, _ := namedReturn(10.8, 5.6)
	fmt.Printf("Area %f", daerah)
}
