package main

import "fmt"

/*
1. Variadic functions can be called with any number of trailing arguments.
2. Variadic functions can be called in the usual way with individual arguments.
3. If you already have multiple args in a slice, apply 
	them to a variadic function using func(slice...) like this.
*/


func sum(nums ...int) {
	fmt.Print(nums, " ")
	// total := 0 
	var total int = 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}


func main() {
	sum(1, 2)
	sum(1, 2, 3)
	
	nums := []int{1, 2, 3, 4}
	sum(nums...) // ?
}
























