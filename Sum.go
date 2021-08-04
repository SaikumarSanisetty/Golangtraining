package main

import (
	"fmt"
	"sort"
)

func add(num1 []int, num2 []int) []int {
	sum := make([]int, len(num1))
	sort.Sort(sort.Reverse(sort.IntSlice(num2)))
	for i, _ := range num1 {
		sum[i] = num1[i] + num2[i]
	}

	return sum
}

func main() {
	fmt.Println("addition of two slices")
	num1 := []int{1, 2, 3, 4, 5, 6}
	num2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(add(num1, num2))

}
