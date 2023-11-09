package main

import "fmt"

func main() {
	var a bool = true
	var b bool = false

	// && = and
	fmt.Println(a && b) // false
	// || = or
	fmt.Println(a || b) // true
	// ! = not
	fmt.Println(!a) // false

	// Boolean operation program
	var nilaiAKhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAKhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println("Apakah Lulus?", lulus)
}
