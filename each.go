package hotchain

// Each runs a function for each element for its side effects.
func (c Chain[T]) Each(fn func(T)) {
	for _, v := range c {
		fn(v)
	}
}
