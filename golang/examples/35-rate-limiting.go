package main 

import "time"
import "fmt"

/*

Rate limiting is an important mechanism for controlling 
resource utilization and maintaining quality of service. 
Go elegantly supports rate limiting with goroutines, 
channels, and tickers.
*/
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
	limiter := time.Tick(time.Millisecond * 200)

	/*
	By blocking on a receive from the limiter channel 
	before serving each request, we limit ourselves 
	to 1 request every 200 milliseconds.
	*/
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	/*
	We may want to allow short bursts of requests in our 
	rate limiting scheme while preserving the overall rate limit. 
	We can accomplish this by buffering our limiter channel. 
	This burstyLimiter channel will allow bursts of up to 3 events.
	*/
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }

    go func() {
    	for t := range time.Tick(time.Millisecond * 1000) {
    		burstyLimiter <- t
    	}
    }()

    burstyRequests := make(chan int, 5)
    for i:= 1; i <= 5; i++ {
    	burstyRequests <- i
    }
    close(burstyRequests)

    for req := range burstyRequests {
    	<-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}
