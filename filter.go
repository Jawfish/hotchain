package hotchain

// Filter returns a new Chain with the elements that satisfy the predicate function.
func (c Chain[T]) Filter(fn func(T) bool) Chain[T] {
	result := make(Chain[T], len(c))
	j := 0
	for i := 0; i < len(c); i++ {
		if fn(c[i]) {
			result[j] = c[i]
			j++
		}
	}
	return result[:j]
}
