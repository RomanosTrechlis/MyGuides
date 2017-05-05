package main 

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	
	f("direct")

	/*
	To invoke this function in a goroutine, use go f(s). 
	This new goroutine will execute concurrently with the calling one.
	*/
	go f("goroutine")


	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
