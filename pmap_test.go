package hotchain

import (
	"sync/atomic"
	"testing"
)

func TestPmap(t *testing.T) {
	tests := []struct {
		name     string
		chain    Chain[int]
		fn       func(int) int
		expected Chain[int]
		counter  *int32
	}{
		{
			name:     "Square of numbers concurrently",
			chain:    Chain[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			fn:       square,
			expected: Chain[int]{1, 4, 9, 16, 25, 36, 49, 64, 81, 100, 121, 144, 169, 196, 225, 256, 289, 324, 361, 400},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.counter != nil {
				tt.fn = func(x int) int { atomic.AddInt32(tt.counter, 1); return x }
			}
			actual := tt.chain.Pmap(tt.fn)
			checkSlice(t, tt.expected, actual, tt.name)

			if tt.counter != nil {
				if got := atomic.LoadInt32(tt.counter); got != int32(len(tt.chain)) {
					t.Errorf("TestPmap %s, counter = %v, want %v", tt.name, got, len(tt.chain))
				}
			}
		})
	}
}
