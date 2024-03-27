// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"strconv"
	"sync"
)

// создаем свою мапу, способную работать при конкурентности
type mySyncMap struct {
	m  map[string]int
	mu sync.RWMutex
}

func genSyncMap() mySyncMap {
	return mySyncMap{make(map[string]int), sync.RWMutex{}}
}

func main() {
	wg := sync.WaitGroup{}
	mm := genSyncMap()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			mm.mu.Lock()

			id := strconv.Itoa(i)
			mm.m[id] = i

			mm.mu.Unlock()
		}(i)
	}

	wg.Wait() // ожидаем выполения горутин

	fmt.Println(mm.m)
}
