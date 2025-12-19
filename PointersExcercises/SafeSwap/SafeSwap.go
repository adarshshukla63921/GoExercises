package main

import "fmt"

func safeSwap(a *int, b *int) {
	if a == nil || b == nil {
		return
	}
	*a , *b = *b , *a
}
func main() {
	x := 10
	y := 20
	safeSwap(&x, &y)

	fmt.Println("x:", x)
	fmt.Println("y:", y)
}