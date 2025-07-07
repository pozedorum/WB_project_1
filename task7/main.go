package main

import (
	"fmt"
	"sync"
)

type myMap struct {
	mx sync.Mutex
	m  map[int]string
}

func (m *myMap) Put(key int, value string) {
	m.mx.Lock()
	m.m[key] = value
	m.mx.Unlock()
}

func (m *myMap) Get(key int) (string, bool) {
	m.mx.Lock()
	value, ok := m.m[key]
	m.mx.Unlock()
	return value, ok
}

func main() {
	m := &myMap{m: make(map[int]string)}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(2)

		go func(i int) {
			m.Put(i, fmt.Sprintf("value_%d", i))
			wg.Done()
		}(i)

		go func(i int) {
			m.Put(i, fmt.Sprintf("value_%d", i+1))
			wg.Done()
		}(i)
	}
	for key := 0; key < 100; key++ {
		value, ok := m.Get(key)
		if !ok {
			fmt.Printf("key %d not found\n", key)
		} else {
			fmt.Printf("%d %s\n", key, value)
		}

	}
	wg.Wait()
}
