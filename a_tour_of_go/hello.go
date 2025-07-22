package main

import (
	"fmt"
	"time"
)

// func swap(x, y int) (int, int) {
// 	return y, x
// }

func main() {
	fmt.Println("When is Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("today")

	}
}
