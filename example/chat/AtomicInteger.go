// AtomicInteger
package main

import (
	"sync"
)

type AtomicInteger struct {
	sync.Mutex
	value int
}

func NewAtomicInteger() AtomicInteger {
	result := AtomicInteger{value: 0}
	return result
}

func (i *AtomicInteger) AddAndGet() int {
	i.Lock()
	i.value += 1
	result := i.value
	i.Unlock()
	return result
}

func (i *AtomicInteger) GetValue() int {
	i.Lock()
	result := i.value
	i.Unlock()
	return result
}
