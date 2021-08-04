package main

import (
	"fmt"
	"sync"
)

const xthreads = 10 // Total number of threads to use

func doSomething(a int) {
	fmt.Println("My job is", a)
	return
}

func main() {
	var ch = make(chan int, 50)
	var wg sync.WaitGroup

	// This starts xthreads number of goroutines
	wg.Add(xthreads)
	for i := 0; i < xthreads; i++ {
		go func() {
			for {
				a, ok := <-ch
				if !ok { // if there is nothing to do and the channel has been closed end the goroutine
					wg.Done()
					return
				}
				doSomething(a)
			}
		}()
	}

	for i := 0; i < 50; i++ {
		ch <- i
	}

	close(ch)
	wg.Wait()
}
