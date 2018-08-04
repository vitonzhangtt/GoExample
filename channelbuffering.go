package main

import "fmt"

/*
1. By default channels are unbuffered, meaning that 
	they will only accept sends (chan <-) if there is a corresponding 
	receive (<- chan) ready to receive the sent value.

2. Buffered channels accept a limited number of values without 
	a corresponding receiver for those values.

3.

*/

func main() {

// Here we make a channel of strings buffering up to 2 values.
	messages := make(chan string, 2)
	
// Because this channel is buffered, we can send these values 
	into the channel without a corresponding concurrent receive.
	
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}


























