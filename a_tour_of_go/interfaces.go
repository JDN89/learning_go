package main

import (
	"fmt"
	"strconv"
)

type Test interface {
	Yolo()
}

type TestStruct struct {
	x, y string
}

func (t *TestStruct) Yolo() {
	fmt.Println(t.x, t.y)
}

// implicit interface we pass here
func DoSometing(t Test) {
	t.Yolo()
}

type Sender interface {
	Send(message string)
}

type RealSender struct{}

type FakeSender struct {
	Messages []string
}

func (f *FakeSender) Send(message string) {
	f.Messages = append(f.Messages, message)
}

// you can pass whichever stuct implements the Sender interface
func Notify(sender Sender) {
	sender.Send("test")
}

func main() {
	// instantiate the struct
	t := TestStruct{"Jan", "De Niels"}
	// we pass the struct that implicitly implements the interfrace
	DoSometing(&t)

	fake := &FakeSender{}
	Notify(fake)
}
