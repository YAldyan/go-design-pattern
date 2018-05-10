package Singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()

	if counter1 == nil {
		t.Error("Expected pointer to singleton after calling is not null")
	}

	expectedCounter1 := counter1

	currentCounter := counter1.AddOne()
	if currentCounter != 1 {
		t.Error("After Calling Once, the counter must be 1, instead of %d\n", currentCounter)
	}

	counter2 := GetInstance()

	if counter2 != expectedCounter1 {
		t.Error("Expected same Instance but it got different instance")
	}

	currentCounter2 := counter2.AddOne()
	if currentCounter2 != 2 {
		t.Error("After Calling AddOne currcnt counter must be 2, instead of %d\n", currentCounter2)
	}
}
