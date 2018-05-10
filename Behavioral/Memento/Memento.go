package Memento

import (
	"fmt"
)

type State struct {
	Description string
}

type memento struct {
	state State
}

type originator struct {
	state State
}

func (o *originator) NewMemento() memento {
	return memento{}
}
func (o *originator) ExtractAndStoreState(m memento) {
	//Does nothing
}

type careTaker struct {
	mementoList []memento
}

func (c *careTaker) Add(m memento) {
	//Does nothing
}

func (c *careTaker) Memento(i int) (memento, error) {
	return memento{}, fmt.Errorf("Not implemented yet")
}

/*Another Example Kombinasi Command Pattern & Facade Pattern dengan Memento Pattern*/

/*
	States are Volumn and Mute


	Logically.

	Command pattern to encapsulate a set of different types of states (those that implement a Command interface)
	Command pattern akan meng-enkapsulasi sekumpulan tipe states yang berbeda-beda, states akan mengextens Command Interface

	provide a small facade to automate the insertion in the caretaker object
	Menyediakan sebuah Facade kecil untuk insert otomatis Memento ke Object Caretaker
*/

type Command interface {
	GetValue() interface{}
}

type Volume byte

func (v Volume) GetValue() interface{} {
	return v
}

type Mute bool

func (m Mute) GetValue() interface{} {
	return m
}

/*
	Memento menyimpan pointer States
	*Volume
	*Mute
*/
type MementoX struct {
	memento Command
}

/*
	Originator untuk membuat Memento
*/
type originatorX struct {
	Command Command
}

func (o *originatorX) NewMementoX() MementoX {
	return MementoX{memento: o.Command}
}
func (o *originatorX) ExtractAndStoreCommandX(m MementoX) {
	o.Command = m.memento
}

/*
	Caretaker Stack of Memento
*/
type careTaker struct {
	mementoList []Memento
}

func (c *careTaker) Add(m Memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *careTaker) Pop() Memento {
	if len(c.mementoStack) > 0 {
		tempMemento := c.mementoStack[len(c.mementoStack)-1]
		c.mementoStack = c.mementoStack[0 : len(c.mementoStack)-1]
		return tempMemento
	}
	return Memento{}
}

type MementoFacade struct {
	originator originator
	careTaker  careTaker
}

func (m *MementoFacade) SaveSettings(s Command) {
	m.originator.Command = s
	m.careTaker.Add(m.originator.NewMemento())
}
func (m *MementoFacade) RestoreSettings(i int) Command {
	m.originator.ExtractAndStoreCommand(m.careTaker.Memento(i))
	return m.originator.Command
}

func assertAndPrint(c Command) {
	switch cast := c.(type) {
	case Volume:
		fmt.Printf("Volume:\t%d\n", cast)
	case Mute:
		fmt.Printf("Mute:\t%t\n", cast)
	}
}
