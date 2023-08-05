package hotchain

import "testing"

func TestReduce(t *testing.T) {
	tests := []struct {
		name     string
		chain    Chain[int]
		fn       func(int, int) int
		initial  int
		expected int
	}{
		{
			name:     "Sum of numbers",
			chain:    Chain[int]{1, 2, 3, 4, 5},
			fn:       add,
			initial:  0,
			expected: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.chain.Reduce(tt.fn, tt.initial)
			check(t, tt.expected, actual)
		})
	}
}
