package main

import "fmt"

func simpleFunction() {
	fmt.Println("Simple funtion")

}
func add(a, b int) (result int) {
	result = a + b
	return
}
func multiply(a, b int) (result int) {
	result = a * b
	return
}
func div(a, b int) (result int) {
	result = a / b
	return
}

func main() {
	fmt.Println("We are learning funtions in golang")
	simpleFunction()

	ans := add(3, 4)
	fmt.Println("add of the two number", ans)

	data := multiply(3, 4)
	fmt.Println("add of the two number", data)

	total := div(4, 3)
	fmt.Println("add of the two number", total)

}
