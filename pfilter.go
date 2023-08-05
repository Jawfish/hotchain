package hotchain

import "sync"

// Pfilter returns a new Chain with the elements that satisfy the predicate function concurrently.
func (c Chain[T]) Pfilter(fn func(T) bool) Chain[T] {
	results := make([]struct {
		val  T
		pass bool
	}, len(c))
	var wg sync.WaitGroup
	wg.Add(len(c))

	for i, v := range c {
		go func(i int, v T) {
			defer wg.Done()
			if fn(v) {
				results[i] = struct {
					val  T
					pass bool
				}{val: v, pass: true}
			}
		}(i, v)
	}

	wg.Wait()

	filtered := Chain[T]{}
	for _, result := range results {
		if result.pass {
			filtered = append(filtered, result.val)
		}
	}

	return filtered
}
