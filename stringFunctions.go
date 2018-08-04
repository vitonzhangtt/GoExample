package main

/*
1. https://medium.com/golangspec/import-declarations-in-go-8de0fd3ae8ff
*/
import s "strings"
import "fmt"

// We alias fmt.Println to a shorter name as we’ll use it a lot below.
var p = fmt.Println

func main() {
	p("Contains:	", s.Contains("test", "es"))
	p("Count:	", s.Count("test", "t"))
	p("HasPrefix:	", s.HasPrefix("test", "te"))
	p("HasSuffix:	", s.HasSuffix("test", "st"))
	p("Index:	", s.Index("test", "e"))
	p("Join:	", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:	", s.Repeat("a", 5))
	p("Replace:	", s.Replace("foo", "o", "0", -1))
	p("Replace:	", s.Replace("foo", "o", "0", 1))
	p("Split:	", s.Split("a-b-c-d-e", "-"))
	p("ToLower:	", s.ToLower("TEST"))
	p("ToUpper:	", s.ToUpper("test"))
	p()
/*
2. the mechanisms for getting the length of a string in bytes and getting a byte by index.
*/

// strings, bytes, runes and characters in Go
// https://blog.golang.org/strings

	p("Len:	", len("hello"))
	p("Char:", "hello"[1])

}

















