package main

import "os"
import "fmt"

/*
1. Link
https://gobyexample.com/command-line-arguments
*/

/*
2. Command-line arguments are a common way to parameterize execution 
	of programs.
*/
func main() {

/*
3. os.Args provides access to raw command-line arguments. 
*/
	argsWithProg := os.Args
/*
4. Note that the first value in this slice is the path to the program, 
	and os.Args[1:] holds the arguments to the program.
*/
	argsWithoutProg := os.Args[1:]
	
/*
5. You can get individual args with normal indexing.
*/
	arg := os.Args[3]
	
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

/*
6. To experiment with command-line arguments it’s best to 
	build a binary with go build first.
*/

/*
$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
*/


