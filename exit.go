package main

import "fmt"
import "os"

/*
1. Link
https://gobyexample.com/exit
*/

/*
2. Use os.Exit to immediately exit with a given status.
*/

func main() {

/*
3. defers will not be run when using os.Exit, so this 
	fmt.Println will never be called.
*/
	defer fmt.Println("!")
	
/*
4. Exit with status 3.
*/

/*
5. Note that unlike e.g. C, Go does not use an integer return 
	value from main to indicate exit status. If you’d like 
	to exit with a non-zero status you should use os.Exit.
*/
	os.Exit(3)
}

/*
6. 
$ go run exit.go

If you run exit.go using go run, the exit will be picked up 
	by go and printed.

*/

/*
7. 
$ go build exit.go
$ ./exit
$ echo $?

By building and executing a binary you can see the status 
	in the terminal.
*/







