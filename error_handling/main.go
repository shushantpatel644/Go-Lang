package main

import "fmt"

func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "denominatormust not be zero"
	}
	return a / b, "nil"
}
func main() {
	fmt.Println("Started error handling in the funtion")
	// ans ,err := divide(10,0)
	// if err != nil{
	// 	fmt.Println(("errror handling"))
	// }
	// fmt.Println("Division of two num is ",ans)

	ans, _ := divide(10, 2)
	fmt.Println("Division of two num is", ans)
}
