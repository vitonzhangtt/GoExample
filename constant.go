package main

import "fmt"
import "math"

/*
1. const declares a constant value.
2. A const statement can appear anywhere a var statement can.
3. Go supports constants of character, string, boolean, and numeric values.
4. A numeric constant has no type until it’s given one, such as by an explicit cast.
*/

const s string = "constant"

func main() {
	fmt.Println(s)
	
	const n = 5000000
	
	const d = 3e20 / n
	fmt.Println(d)
	
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}	



