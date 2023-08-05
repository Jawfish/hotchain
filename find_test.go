package hotchain

import (
	"reflect"
	"testing"
)

func TestFind(t *testing.T) {
	tests := []struct {
		name     string
		chain    Chain[int]
		fn       func(int) bool
		expected Chain[int]
	}{
		{
			name:     "Find first even number",
			chain:    Chain[int]{1, 2, 3, 4, 5},
			fn:       even,
			expected: Chain[int]{2},
		},
		{
			name:     "Find first number greater than three",
			chain:    Chain[int]{1, 2, 3, 4, 5},
			fn:       func(x int) bool { return x > 3 },
			expected: Chain[int]{4},
		},
		{
			name:     "Find non-existing condition returns empty chain",
			chain:    Chain[int]{1, 2, 3, 4, 5},
			fn:       func(x int) bool { return x > 10 },
			expected: Chain[int]{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.chain.Find(tt.fn)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected: %v, but got: %v", tt.expected, actual)
			}
		})
	}
}
