package main

import "fmt"

func swap(a, b *int) {
	a, b = b, a
}

func main() {
	x := 10
	y := 20
	swap(&x, &y)
	fmt.Println(x, y)
}
