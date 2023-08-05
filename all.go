package hotchain

// All returns true if all elements satisfy the predicate function.
func (c Chain[T]) All(fn func(T) bool) bool {
	for _, v := range c {
		if !fn(v) {
			return false
		}
	}
	return true
}
