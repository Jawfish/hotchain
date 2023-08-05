package hotchain

import "testing"

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		chain    Chain[int]
		fn       func(int) int
		expected Chain[int]
	}{
		{
			name:     "Square of numbers",
			chain:    Chain[int]{1, 2, 3, 4, 5},
			fn:       func(x int) int { return x * x },
			expected: Chain[int]{1, 4, 9, 16, 25},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.chain.Map(tt.fn)
			for i, val := range result {
				if val != tt.expected[i] {
					t.Errorf("At index %d: expected %d, got %d", i, tt.expected[i], val)
				}
			}
		})
	}
}
