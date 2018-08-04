package main

import (
	"sort"
	"fmt"
)

/*
1. Sometimes we’ll want to sort a collection by something other than its natural order.  
*/

/*
2. In order to sort by a custom function in Go, we need a corresponding type. 
	Here we’ve created a byLength type that is just an alias for the builtin []string type.
*/
type byLength []string

/*
3. We implement sort.Interface - Len, Less, and Swap - on our type so we can 
	use the sort package’s generic Sort function. 
*/
func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
/*
4. With all of this in place, we can now implement our custom sort by 
	casting the original fruits slice to byLength, and then use 
	sort.Sort on that typed slice.
*/
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}

/*
5. By following this same pattern of creating a custom type, implementing 
	the three Interface methods on that type, and then calling sort.Sort 
	on a collection of that custom type, we can sort Go slices by arbitrary functions.

*/
























































