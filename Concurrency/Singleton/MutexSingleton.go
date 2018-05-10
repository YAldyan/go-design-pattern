package Singleton

import (
	"sync"
)

type singletonX struct {
	count int

	/*
		Read Lock
		Write Lock
	*/
	sync.RWMutex
}

var instanceX singletonX

func GetInstanceX() *singletonX {
	return &instanceX
}

func (s *singletonX) AddOne() {
	s.Lock()
	defer s.Unlock()
	s.count++
}

func (s *singletonX) GetCount() int {
	s.RLock()
	defer s.RUnlock()
	return s.count
}
