package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "Guest", "enter your name")
	flag.Parse()
	fmt.Println("hello", *name)
}
