package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string, 1)

	// menerima data
	go func(ch chan<- string) {
		ch <- "Hello World!"
		println("Finishing goroutine")
	}(channel)

	time.Sleep(time.Second)
	message := <-channel
	fmt.Println(message)
}

// mengirim data
func receivingCh(ch <-chan string) {
	msg := <-ch
	println(msg)
}
