package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Wahyu"
	names[1] = "Setiawan"
	names[2] = "Usman"

	fmt.Println(names[0], names[1], names[2])

	// var values = [3]int{
	// 	90,
	// 	80,
	// 	70,
	// }
	// fmt.Println(values)
	// fmt.Println(values[0], values[1], values[2])

	var values = [...]int{
		90,
		80,
		70,
		60,
	}
	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 100
	fmt.Println(values)
}
