package main

import "fmt"

func KonversiNilai(nilai int) string {
	// your code here
	var study string

	if nilai >= 80 && nilai <= 100 {
		study = ("A")
	} else if nilai >= 65 && nilai <= 79 {
		study = ("B+")
	} else if nilai >= 50 && nilai <= 64 {
		study = ("B")
	} else if nilai >= 35 && nilai <= 49 {
		study = ("C")
	} else {
		study = ("D")
	}
	return study
}

func main() {
	var nilai int = 80

	fmt.Println(KonversiNilai(nilai))
}
