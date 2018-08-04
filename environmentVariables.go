package main


import "os"
import "strings"
import "fmt"

/*
1. Link
https://gobyexample.com/environment-variables
*/

/*
2. http://www.12factor.net/config

*/
func main() {

/*
3. To set a key/value pair, use os.Setenv. To get a value for a key, 
	use os.Getenv. This will return an empty string if the key 
	isn’t present in the environment.
*/
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	
/*
4. Use os.Environ to list all key/value pairs in the environment. 
	This returns a slice of strings in the form KEY=value. 
	You can strings.Split them to get the key and value. Here 
	we print all the keys.
*/
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
	//	fmt.Println(pair[0])
		fmt.Printf("%s:%s \n", pair[0], pair[1])
	}
}

/*
5. 
$ go run environment-variables.go
Running the program shows that we pick up the value for FOO 
	that we set in the program, but that BAR is empty.
*/


/*
6. 
$ BAR=2 go run environment-variables.go
If we set BAR in the environment first, the running program 
	picks that value up.
*/


















