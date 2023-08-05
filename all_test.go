package hotchain

import "testing"

func TestAll(t *testing.T) {
	tests := []struct {
		name     string
		chain    Chain[int]
		fn       func(int) bool
		expected bool
	}{
		{
			name:     "All numbers are even",
			chain:    Chain[int]{2, 4, 6, 8, 10},
			fn:       even,
			expected: true,
		},
		{
			name:     "All numbers are greater than 10",
			chain:    Chain[int]{2, 4, 6, 8, 10},
			fn:       func(x int) bool { return x > 10 },
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.chain.All(tt.fn)
			check(t, tt.expected, actual)
		})
	}
}
