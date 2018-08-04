package main

import "fmt"

/*
1. A goroutine is a lightweight thread of execution.
2.
*/

func f(from string) {
	for i := 0; i < 9; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")
/*
To invoke this function in a goroutine, use go f(s). This new goroutine 
will execute concurrently with the calling one.
*/	
	go f("goroutine")
	

	go func(msg string) {
		fmt.Println(msg)
	}("going")
/*
This Scanln requires we press a key before the program exits.
*/
	fmt.Scanln()
	fmt.Println("done")

}




























