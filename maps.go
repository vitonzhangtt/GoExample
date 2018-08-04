package main

import "fmt"

/*
1. Maps are Go’s built-in associative data type 
(sometimes called hashes or dicts in other languages).

2. To create an empty map, use the builtin make: make(map[key-type]val-type).
*/

func main() {

	m := make(map[string]int)

// Set key/value pairs using typical name[key] = val syntax.		
	m["k1"] = 7
	m["k2"] = 13
	
	fmt.Println("map:", m)

// Get a value for a key with name[key].
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

// The builtin len returns the number of key/value pairs when called on a map.	
	fmt.Println("len: ", len(m))

// The builtin delete removes key/value pairs from a map.	
	delete(m, "k2")
	fmt.Println("map: ", m)

// "_" blank identifier 	
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

// You can also declare and initialize a new map 
// in the same line with this syntax.
	n := map[string]int{"foo":1, "bar":2}
	fmt.Println("map:", n)
}

























