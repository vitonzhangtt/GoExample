package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
1. Link
https://gobyexample.com/line-filters
*/

/*
2. A line filter is a common type of program that reads input on stdin, 
	processes it, and then prints some derived result to stdout. 
	grep and sed are common line filters.
*/

func main() {
	
/*
3. Wrapping the unbuffered os.Stdin with a buffered scanner gives 
	us a convenient Scan method that advances the scanner 
	to the next token; which is the next line in the default scanner.
*/
	scanner := bufio.NewScanner(os.Stdin)
	
// ??? for-clause ???
	for scanner.Scan() {
/*
4. Text returns the current token, here the next line, from the input.
*/
		ucl := strings.ToUpper(scanner.Text())
/*
5. Write out the uppercased line.
*/
		fmt.Println(ucl)
	}
	
/*
6. Check for errors during Scan. End of file is expected and 
	not reported by Scan as an error.
*/
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}




