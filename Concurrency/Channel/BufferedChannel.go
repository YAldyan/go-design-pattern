package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	channel := make(chan string, 1)

	go func() {
		channel <- "Hello World! 1"
		channel <- "Hello World! 2"
		println("Finishing goroutine")
	}()

	time.Sleep(time.Second)

	message := <-channel
	fmt.Println(message)

	tidakBlocking()
}

func tidakBlocking() {
	runtime.GOMAXPROCS(2)

	messages := make(chan int, 2)

	go func() {
		for {
			i := <-messages
			fmt.Println("receive data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		messages <- i
	}
}
