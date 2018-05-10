package main

import "fmt"

func main() {
	channel := make(chan string)

	/*
		main program tidak akan selesai lebih dahulu
		dikarenakan goroutine dengan channel akan
		blocking sampai nilai dari channel di serahkan
		ke variabel d main program.
	*/
	go func() {
		channel <- "Hello World!"
	}()

	message := <-channel
	fmt.Println(message)
}

func withWaitGroup() {
	channel := make(chan string)

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go func() {
		channel <- "Hello World!"
		println("Finishing goroutine")
		waitGroup.Done()
	}()

	time.Sleep(time.Second)

	message := <-channel
	fmt.Println(message)
	waitGroup.Wait()
}
