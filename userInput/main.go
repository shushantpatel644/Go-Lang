package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter your name")
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Hello my name is", name)

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello my name is", name)
}
