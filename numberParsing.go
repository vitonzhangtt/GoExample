package main

import "strconv"
import "fmt"

/*
1. Link
https://gobyexample.com/number-parsing

2. The built-in package strconv provides the number parsing.
*/


func main() {

/*
3. With ParseFloat, this 64 tells how many bits of precision to parse.
*/
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

/*
4. For ParseInt, the 0 means infer the base from the string. 64 requires that the result fit in 64 bits.
*/
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)
	
/*
5. ParseInt will recognize hex-formatted numbers.
*/
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)
	
/*
6. A ParseUint is also available.
*/
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

/*
7. Atoi is a convenience function for basic base-10 int parsing.
*/	
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

/*
8. Parse functions return an error on bad input.
*/
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}




