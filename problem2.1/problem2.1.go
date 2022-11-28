package main

import "fmt"

func FaktorBilangan(n int) string {
	// your code here
	var res string
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			res += fmt.Sprintln(i)
		}
	}
	return res
}

func main() {
	var number int
	fmt.Scanf("%d", &number)
	fmt.Println(FaktorBilangan(number))
}
