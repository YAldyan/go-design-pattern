package Observer

import (
	"fmt"
)

type Observer interface {

	/*
		Notify method that accepts a string
		type that will contain the message
		to spread
	*/
	Notify(string)
}

type TestObserver struct {
	ID      int
	Message string
}

func (p *TestObserver) Notify(m string) {
	fmt.Printf("Observer %d: message '%s' received \n", p.ID, m)
	p.Message = m
}

/*
	Kelas yang akan mendapatkan Update dari Observer
*/
type Publisher struct {
	ObserversList []Observer
}

// subscribe a new observer to the publisher
func (s *Publisher) AddObserver(o Observer) {
	s.ObserversList = append(s.ObserversList, o)
}

// unsubscribe an observer
func (s *Publisher) RemoveObserver(o Observer) {
	var indexToRemove int
	for i, observer := range s.ObserversList {
		if observer == o {
			indexToRemove = i
			break
		}
	}

	/*
		The way to remove indexes on a slice in Go is a bit tricky:
		1. 	First, we need to use slice indexing to return a new slice containing every object from the beginning of the slice
			to the index we want to remove (not included).
		2. 	Then, we get another slice from the index we want to remove (not included) to the last object in the slice
		3. 	Finally, we join the previous two new slices into a new one (the append function)

		For example, in a list from 1 to 10 in which we want to remove the number 5, we have to create a new slice, joining a
		slice from 1 to 4 and a slice from 6 to 10.
	*/
	s.ObserversList = append(s.ObserversList[:indexToRemove], s.ObserversList[indexToRemove+1:]...)
}

// method with a string that acts as the message we want to spread between all observers
func (s *Publisher) NotifyObservers(m string) {
	fmt.Printf("Publisher received message '%s' to notify observers\n", m)
	for _, observer := range s.ObserversList {
		observer.Notify(m)
	}
}
