package main

import "fmt"

/*
1. Go supports methods defined on struct types. 
*/

type rect struct {
	width, height int
}

/*
2. This area method has a receiver type of *rect.
*/
func (r *rect) area() int {
	return r.width * r.height
}
/*
3. Methods can be defined for either pointer or value receiver types. 
	Here’s an example of a value receiver.
*/
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

/*
4. Go automatically handles conversion between values and pointers for method calls.
*/
func main() {
	r := rect{width: 10, height:5}
	
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())
	
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
	
}
