package hotchain

// Map returns a new Chain with the results of applying a function to each element.
func (c Chain[T]) Map(fn func(T) T) Chain[T] {
	result := make(Chain[T], len(c))
	for i, v := range c {
		result[i] = fn(v)
	}
	return result
}
