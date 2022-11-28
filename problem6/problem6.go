package main

import (
	"fmt"
	"strconv"
)

func FullPrima(n int) bool {
	// write your code
	var totalSemua bool = true
	var total bool = true
	var nStr string

	total = PrimeNumber(n)

	if total {
		nStr = strconv.Itoa(n)

		for i := 0; i < len(nStr); i++ {
			keseAngka, err := strconv.Atoi(string(nStr[i]))

			if err != nil {
				fmt.Println("string error", string(nStr[i]), "Error =", err)
			}
			if !PrimeNumber(keseAngka) {
				totalSemua = false
			}
		}
	}
	return totalSemua
}
func PrimeNumber(number int) bool {
	// your code here
	var prima int
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			prima++
		}
	}
	if prima == 2 {
		return true
	}
	return false
}

func main() {
	fmt.Println(FullPrima(2))  // true
	fmt.Println(FullPrima(3))  // true
	fmt.Println(FullPrima(11)) // false
	fmt.Println(FullPrima(13)) // false
	fmt.Println(FullPrima(23)) // true
	fmt.Println(FullPrima(29)) // false
	fmt.Println(FullPrima(37)) // true
	fmt.Println(FullPrima(41)) // false
	fmt.Println(FullPrima(43)) // false
	fmt.Println(FullPrima(53)) // true
}
