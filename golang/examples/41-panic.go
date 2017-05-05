package main 

import "os"

func main() {
	panic("a problem")

	/*
	Note that unlike some languages which use exceptions for 
	handling of many errors, in Go it is idiomatic to use 
	error-indicating return values wherever possible.
	*/
	_, err := os.Create("./error")
	if err != nil {
		panic(err)
	}
}
