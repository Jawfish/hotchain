package hotchain

import (
	"runtime"
	"sync"
)

// Pmap returns a new Chain with the results of applying a function to each element concurrently.
func (c Chain[T]) Pmap(fn func(T) T) Chain[T] {
	result := make(Chain[T], len(c))

	workers := runtime.NumCPU()
	size := len(c) / workers
	if len(c)%workers != 0 {
		size++
	}

	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			start := i * size
			end := start + size
			if end > len(c) {
				end = len(c)
			}
			for j := start; j < end; j++ {
				result[j] = fn(c[j])
			}
		}(i)
	}

	wg.Wait()

	return result
}
