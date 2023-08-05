package hotchain

// Find returns a new Chain with the first element that satisfies the predicate function, or an empty Chain if none do.
func (c Chain[T]) Find(fn func(T) bool) Chain[T] {
	for _, v := range c {
		if fn(v) {
			return Chain[T]{v}
		}
	}
	return Chain[T]{}
}
