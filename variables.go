package main

import "fmt"

/*
1. var declares 1 or more variables.
2. Go will infer the type of initialized variables.
3. Variables declared without a corresponding initialization 
are zero-valued. For example, the zero value for an int is 0.
4. The := syntax is shorthand for declaring and initializing 
a variable, e.g. for var f string = "short" in this case.

*/
func main() {
	var a = "initial"
	fmt.Println(a)
	
	var b, c int = 1, 2 
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)
}



