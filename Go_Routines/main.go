package main

import (
	"fmt"
	"time"
)

func greet(text string) {
	fmt.Println("Normal Func: ", text)
}

// will run in goroutine
func slowGreet(text string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("From Slow Func: ", text)
	doneChan <- true //send data to the channel back to where it start

}
func main() {
	done := make(chan bool)
	greet("hello")

	go slowGreet("hi", done)
	greet("hey")
	<-done //read data from channel

}
