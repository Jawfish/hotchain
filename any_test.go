package hotchain

import "testing"

func TestAny(t *testing.T) {
	tests := []struct {
		name     string
		chain    Chain[int]
		fn       func(int) bool
		expected bool
	}{
		{
			name:     "Any even number",
			chain:    Chain[int]{1, 2, 3, 4, 5},
			fn:       even,
			expected: true,
		},
		{
			name:     "Any number greater than ten",
			chain:    Chain[int]{1, 2, 3, 4, 5},
			fn:       func(x int) bool { return x > 10 },
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			any := tt.chain.Any(tt.fn)
			check(t, tt.expected, any)
		})
	}
}
