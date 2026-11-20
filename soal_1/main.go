package main

import "fmt"

// Prosedur tukar: menukar nilai a dan b
func tukar(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// Prosedur urutTiga: mengurutkan a, b, c dari terkecil ke terbesar
func urutTiga(a *int, b *int, c *int) {
	if *a > *b {
		tukar(a, b)
	}
	if *b > *c {
		tukar(b, c)
	}
	if *a > *b {
		tukar(a, b)
	}
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Println(a, b, c)

	urutTiga(&a, &b, &c)

	fmt.Println(a, b, c)
}