package main

import "fmt"
import "time"

/*
1. Link
https://gobyexample.com/epoch
*/

/*
2. Unix time
http://en.wikipedia.org/wiki/Unix_time
*/
func main() {
	
/*
3. Use time.Now with Unix or UnixNano to get elapsed time since 
	the Unix epoch in seconds or nanoseconds, respectively.
*/
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)
/*
4. You can also convert integer seconds or nanoseconds since 
	the epoch into the corresponding time.
*/	
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}



































