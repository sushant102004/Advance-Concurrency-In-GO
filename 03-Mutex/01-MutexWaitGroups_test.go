package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestState(t *testing.T) {
	wg := &sync.WaitGroup{}
	mx := &sync.Mutex{}
	state := State{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			mx.Lock()
			state.count = i + 1
			mx.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(state)
}
