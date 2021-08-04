package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("logical processor:", runtime.NumCPU())
	fmt.Println("operating system:", runtime.GOOS)
	fmt.Println("system architecture:", runtime.GOARCH)
	fmt.Println("maximum processes:", runtime.GOMAXPROCS(12))
	wg.Add(2)
	go EvenNumbers()
	go OddNumbers()
	fmt.Println("main routine invoked")
	wg.Wait()
}
func EvenNumbers() {
	defer wg.Done()
	fmt.Println("even numbers 1 to 20")
	num := 20
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
}
func OddNumbers() {
	defer wg.Done()
	fmt.Println("Odd numbers 1 to 20")
	num := 20
	for i := 1; i <= num; i++ {
		if i%2 != 0 {
			fmt.Print(i, " ")
		}
	}
}
