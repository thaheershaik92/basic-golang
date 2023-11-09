package main

import "fmt"

func main() {
	type NoKTP string

	var ktpWahyu NoKTP = "1234567890"

	var contoh string = "0987654321"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpWahyu)
	fmt.Println(contoh)
	fmt.Println(contohKtp)
}
