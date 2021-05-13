package main

import "fmt"
import "math"

/*
1. Interfaces are named collections of method signatures.

http://jordanorelli.tumblr.com/post/32665860244/how-to-use-interfaces-in-go
*/

// interface
type geometry interface {
	area() float64
	perim() float64
}

// struct
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

/*
2. To implement an interface in Go, we just need to 
	implement all the methods in the interface.
*/
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

/*
3. If a variable has an interface type, then we can call 
	methods that are in the named interface.
*/

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}


func main() {
	r := rect{width: 3, height:4}
	c := circle{radius: 5}
	
	measure(r)
	measure(c)
}


/*
Reference
====

1. Interfaces in Go
https://medium.com/rungo/interfaces-in-go-ab1601159b3a

*/

























