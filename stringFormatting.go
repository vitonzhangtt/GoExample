package main

import "fmt"
import "os"

/*
1. Printf
*/

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
// 2. this prints an instance of our point struct.
	fmt.Printf("%v\n", p)

// 3. If the value is a struct, the %+v variant will include 
// the struct’s field names.
	fmt.Printf("%+v\n", p)

/*
4. The %#v variant prints a Go syntax representation of the value, 
	i.e. the source code snippet that would produce that value.
*/	
	fmt.Printf("%#v\n", p)
	
/*
5. To print the type of a value, use %T.
*/
	fmt.Printf("%T\n", p)
	
/*
6. Formatting booleans is straight-forward.
*/
	fmt.Printf("%t\n", true)

/*
7. Use %d for standard, base-10 formatting.
*/
	fmt.Printf("%d\n", 123)
	
/*
8. This prints a binary representation.
*/
	fmt.Printf("%b\n", 14)

/*
9. This prints the character corresponding to the given integer.
*/
	fmt.Printf("%c\n", 33)

/*
10. %x provides hex encoding.
*/
	fmt.Printf("%x\n", 456)
/*
11. There are also several formatting options for floats. 
	For basic decimal formatting use %f.
*/
	fmt.Printf("%f\n", 78.9)
/*
12. %e and %E format the float in (slightly different versions of) scientific notation.
*/
	fmt.Printf("%e\n", 123400000.0)

	fmt.Printf("%E\n", 123400000.0)

/*
13. For basic string printing use %s.
*/
	fmt.Printf("%s\n", "\" string \"")
	
/*
14. To double-quote strings as in Go source, use %q.
*/
	fmt.Printf("%q\n", "\" string \"")
/*
15. As with integers seen earlier, %x renders the string in base-16, 
	with two output characters per byte of input.
*/
	fmt.Printf("%x\n", "hex this")

/*
16. To print a representation of a pointer, use %p.
*/
	fmt.Printf("%p\n", &p)

/*
17. When formatting numbers you will often want to control the width 
	and precision of the resulting figure. To specify the width 
	of an integer, use a number after the % in the verb. By default 
	the result will be right-justified and padded with spaces.
*/
	fmt.Printf("|%6d|%6d|\n", 12, 345)

/*
18. You can also specify the width of printed floats, though usually 
	you’ll also want to restrict the decimal precision at the same time 
	with the width.precision syntax.
*/
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	
/*
19. To left-justify, use the - flag.
*/
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
/*
20. You may also want to control width when formatting strings, 
	especially to ensure that they align in table-like output. 
	For basic right-justified width.
*/
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
/*
21. To left-justify use the - flag as with numbers.
*/
	fmt.Printf("|-%6s|-%6s|\n", "foo", "b")
/*
22. Sprintf formats and returns a string without printing it anywhere.
*/
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

/*
23. You can format+print to io.Writers other than os.Stdout using Fprintf.
*/
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}














































