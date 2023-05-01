package main

import "fmt"

func sum(nums ...int) int {
	tot := 0
	for _, num := range nums {
		tot += num
	}
	return tot
}

func main() {
	array := [...]int{10, 12, 14, 28, 30}
	fmt.Printf("%T\n\n", array)

	// fmt.Println(sum(1, 2, 3))
	// fmt.Println(sum(4, 5, 6, 7))
}
