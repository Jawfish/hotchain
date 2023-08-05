package hotchain

// Reduce returns a single value by applying a function to each element and the previous result.
func (c Chain[T]) Reduce(fn func(T, T) T, initial T) T {
	result := initial
	for _, v := range c {
		result = fn(result, v)
	}
	return result
}
