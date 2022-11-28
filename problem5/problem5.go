package main

import "fmt"

func Pangkat(base, pangkat int) int {
	// your code here
	var total int = 1

	for i := 0; i < pangkat; i++ {
		total = total * base
	}
	return total
}

func main() {
	fmt.Println(Pangkat(2, 3))  // 8
	fmt.Println(Pangkat(5, 3))  // 125
	fmt.Println(Pangkat(10, 2)) // 100
	fmt.Println(Pangkat(2, 5))  // 32
	fmt.Println(Pangkat(7, 3))  // 343
}
