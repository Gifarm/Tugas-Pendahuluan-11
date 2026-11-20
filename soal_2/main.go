package main

import (
	"fmt"
	"math"
)


// Prosedur persegiPanjang
func persegiPanjang(p, l float64, luas *float64, keliling *float64) {
	*luas = p * l
	*keliling = 2 * (p + l)
}

// Prosedur lingkaran
func lingkaran(r float64, luas *float64, keliling *float64) {
	*luas = PI * r * r
	*keliling = 2 * PI * r
}

// Prosedur segitigaSiku
func segitigaSiku(a, t float64, luas *float64, keliling *float64) {
	sisiMiring := math.Sqrt(a*a + t*t)
	*luas = a * t / 2
	*keliling = a + t + sisiMiring
}

func main() {
	var kode string
	fmt.Scan(&kode)

	if kode == "P" {
		var p, l float64
		fmt.Scan(&p, &l)
		var luas, keliling float64
		persegiPanjang(p, l, &luas, &keliling)
		fmt.Printf("Luas = %.2f\n", luas)
		fmt.Printf("Keliling = %.2f\n", keliling)
	} else if kode == "L" {
		var r float64
		fmt.Scan(&r)
		var luas, keliling float64
		lingkaran(r, &luas, &keliling)
		fmt.Printf("Luas = %.2f\n", luas)
		fmt.Printf("Keliling = %.2f\n", keliling)
	} else if kode == "S" {
		var a, t float64
		fmt.Scan(&a, &t)
		var luas, keliling float64
		segitigaSiku(a, t, &luas, &keliling)
		fmt.Printf("Luas = %.2f\n", luas)
		fmt.Printf("Keliling = %.2f\n", keliling)
	} else {
		fmt.Println("Kode bangun tidak valid.")
	}}

const PI = 3.14
