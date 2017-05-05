package main 

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i) // i do not pass a reference

	zeroptr(&i) // with & i take the address of variable i
	fmt.Println("zeroptr:", i) // this changes the value of i

	fmt.Println("pointer:", &i)
}
