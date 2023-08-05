package hotchain

// Any returns true if any element satisfies the predicate function.
func (c Chain[T]) Any(fn func(T) bool) bool {
	for _, v := range c {
		if fn(v) {
			return true
		}
	}
	return false
}
