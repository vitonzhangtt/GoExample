package main

import "time"
import "fmt"

/*
1. Timers are for when you want to do something once in the future - tickers 
	are for when you want to do something repeatedly at regular intervals.
*/

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	
	go func() {
/*
Tickers use a similar mechanism to timers: a channel that is sent values. 
Here we’ll use the range builtin on the channel to iterate over the values 
as they arrive every 500ms.
*/
		for t:= range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
/*
Tickers can be stopped like timers. Once a ticker is stopped it won’t receive 
any more values on its channel. We’ll stop ours after 1600ms.
*/
	time.Sleep(1600 * time.Millisecond) 
	ticker.Stop()
	fmt.Println("Ticker stopped")
}









































