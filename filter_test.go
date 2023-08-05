package hotchain

import "testing"

func TestFilter(t *testing.T) {
	tests := []struct {
		name     string
		chain    Chain[int]
		fn       func(int) bool
		expected Chain[int]
	}{
		{
			name:     "Filter even numbers",
			chain:    Chain[int]{1, 2, 3, 4, 5},
			fn:       func(x int) bool { return x%2 == 0 },
			expected: Chain[int]{2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.chain.Filter(tt.fn)
			for i, val := range result {
				if val != tt.expected[i] {
					t.Errorf("At index %d: expected %d, got %d", i, tt.expected[i], val)
				}
			}
		})
	}
}
