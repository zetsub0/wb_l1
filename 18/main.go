// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	sync.Mutex
}

func main() {
	incrementers := 1000
	wg := sync.WaitGroup{}
	c := Counter{}

	for range incrementers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()
	fmt.Println(c.value)
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.value++
}
