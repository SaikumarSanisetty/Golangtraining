package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
func main() {
	x := 10
	y := 20
	fmt.Printf("Before swap,x=%d and y=%d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After swap,x=%d and y=%d", x, y)
}
