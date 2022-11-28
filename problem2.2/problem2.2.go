package main

import "fmt"

func FaktorBilanganDesc(n int) string {
	// your code here
	var res string
	for i := n; i > 0; i-- {
		if n%i == 0 {
			res += fmt.Sprintln(i)
		}
	}
	return res
}

func main() {
	var number int
	fmt.Scanf("%d", &number)
	fmt.Println(FaktorBilanganDesc(number))
}
