package main

import (
	"fmt"
)

func changeA(a int) int {
	a = 20
	fmt.Println(a)
	return a
}

func pointerB(b *int) {
	*b = 20
}

func copyPointerFuckedUp(c *int) {
	b := 10
	c = &b
	fmt.Printf("c = %d", *c)
}

func copyPointerDerefAndChangeWhatHappens(d *int) {
	b := 10
	*d = b
	fmt.Printf("d = %d \n", *d)
}

func main() {
	// NOTE : ponter chananingaens typoooos
	a := 10
	fmt.Println(a)

	// outer a stays 10
	changeA(a)

	// outer a would be 20
	// a = changeA(a)
	fmt.Println(a)

	fmt.Println("hello")
	b := 10
	fmt.Println(b)
	pointerB(&b)
	fmt.Println(b)

	c := 5
	copyPointerFuckedUp(&c)
	// Because we copied the actual pointer and changed the addres the copy points to in the inner function.
	// our local pointer (copy) in the inner scope gets dropped, we pointed to a new address (the og value in main didn't get changed)
	// this gets dropped when function exits, so inner scope and copy gets dropped
	fmt.Printf("c = %d", c)

	d := 5
	copyPointerDerefAndChangeWhatHappens(&d)
	// Because we copied the actual pointer and changed the addres the copy points to in the inner function.
	// our local pointer (copy) in the inner scope gets dropped, we pointed to a new address (the og value in main didn't get changed)
	// this gets dropped when function exits, so inner scope and copy gets dropped
	fmt.Printf("d = %d \n", d)

	// NOTE the distinction is chaning what the points point to vs changing the pointer itself (copy pointer to funciton) go is COPY by VALUE

	// NOTE Zero vs no value
	// in Go, use a normal value (with its zero value) by default. Use a pointer when you genuinely need to distinguish between "zero" and "not provided".

	// NOTE slicing and dicing, appending fucks up capacity (doesn't this cause weird bugs?)

	loafSlice := make([]int, 3)
	for i := range loafSlice {
		loafSlice[i] = 3
	}
	fmt.Println(loafSlice)
	passSliceAndAppend(loafSlice)
	// NOTE, this causes PROBABLY BUGS for some peole !! the appended values are only visible in the function body. the outer slice that we passed stays unchanged, even tough we passed the header (pointer to underlying arr, size capacity), so the header stays unchanged in the outer !
	// FIX : return the slice or pass pointer to slice
	fmt.Println(loafSlice)
}

func passSliceAndAppend(loafSlice []int) {
	loafSlice = append(loafSlice, 4)
	loafSlice = append(loafSlice, 5)
	fmt.Println(loafSlice)
}
