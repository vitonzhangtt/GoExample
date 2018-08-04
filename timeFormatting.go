package main

import "fmt"
import "time"

/*
1. Link
https://gobyexample.com/time-formatting-parsing
*/

/*
2. Go supports time formatting and parsing via pattern-based layouts.
*/

func main() {
	
	p := fmt.Println
	
/*
3. Here’s a basic example of formatting a time according to RFC3339, 
	using the corresponding layout constant.
*/
	t := time.Now()
	p(t.Format(time.RFC3339))
	
/*
4. Time parsing uses the same layout values as Format.
*/
	t1, e := time.Parse(
		time.RFC3339,
		"2018-07-25T22:08:44+00:00")
	p(t1)
/*
5. Format and Parse use example-based layouts. 
*/	
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)	
	
/*
6. For purely numeric representations you can also use 
	standard string formatting with the extracted components 
	of the time value.

*/
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n", 
		t.Year(), t.Month(), t.Day(), 
		t.Hour(), t.Minute(), t.Second())

/*
7. Parse will return an error on malformed input explaining the 
	parsing problem.
*/
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}





