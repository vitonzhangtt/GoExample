package main

import "fmt"
/*
1. Channels are the pipes that connect concurrent goroutines.
2. You can send values into channels from one goroutine and 
	receive those values into another goroutine.
3. Create a new channel with make(chan val-type). 
	Channels are typed by the values they convey.
4. Send a value into a channel using the channel <- syntax. 
5. The <-channel syntax receives a value from the channel. 
6. By default sends and receives block until both the 
	sender and receiver are ready.

*/

func main() {
	
	messages := make(chan string)

// Send a value into a channel using the channel <- syntax. 
	go func() { messages <- "ping"}()	

// The <-channel syntax receives a value from the channel. 
	msg := <-messages
	fmt.Println(msg)
}




