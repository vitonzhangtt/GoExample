package main

import "time"
import "fmt"


func main() {
	
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

/*
This limiter channel will receive a value every 200 milliseconds. 
This is the regulator in our rate limiting scheme.
*/	
	limiter := time.Tick(200 * time.Millisecond)
	
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)
	
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

/*
Every 200 milliseconds we’ll try to add a new value to burstyLimiter, 
up to its limit of 3.
*/
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}







































