package main

import (
	"fmt"
	"time"
)

func greet(text string, doneChan chan bool) {
	fmt.Println("Normal Func: ", text)
	doneChan <- true
}
func slowGreet(text string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("From Slow Func: ", text)
	doneChan <- true

}
func main() {
	//creating a channel
	done := make(chan bool, 3)
	go greet("hello", done)

	go slowGreet("hi", done)
	go greet("hey", done)
	<-done
	<-done
	<-done

}
