package Singleton

import (
	"fmt"
	"testing"
	"time"
)

func TestStartInstance(t *testing.T) {
	singleton := GetInstance()
	singleton2 := GetInstance()
	n := 5000

	for i := 0; i < n; i++ {
		go singleton.AddOne()
		go singleton2.AddOne()
	}

	fmt.Printf("Before loop, current count 1 is %d\n", singleton.GetCount())
	fmt.Printf("Before loop, current count 2 is %d\n", singleton2.GetCount())

	var val int

	for val != n*2 {
		val = singleton.GetCount()
		fmt.Printf("Nilai Val in Looping are %d\n", val)
		time.Sleep(10 * time.Millisecond)
	}

	singleton.Stop()
}
