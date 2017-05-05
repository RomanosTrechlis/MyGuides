package main 

import "fmt"
import "time"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	/*
	Here’s the worker goroutine. It repeatedly receives from jobs 
	with j, more := <-jobs. In this special 2-value form of receive, 
	the more value will be false if jobs has been closed and all values 
	in the channel have already been received. 
	We use this to notify on done when we’ve worked all our jobs.
	*/
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true 
				return
			}
		}
	}()

	for j := 1; j < 3; j++ {
		jobs <-j
		fmt.Println("sent job", j)
		time.Sleep(time.Second * 1) // sent jobs very fast and the result doesn't show 
	}

	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
