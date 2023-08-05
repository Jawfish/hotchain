package hotchain

import (
	"sync/atomic"
	"testing"
)

func TestPfilter(t *testing.T) {
	tests := []struct {
		name     string
		chain    Chain[int]
		fn       func(int) bool
		expected Chain[int]
		counter  *int32
	}{
		{
			name:     "Concurrently filter even numbers",
			chain:    Chain[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
			fn:       even,
			expected: Chain[int]{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.counter != nil {
				tt.fn = func(x int) bool { atomic.AddInt32(tt.counter, 1); return even(x) }
			}
			actual := tt.chain.Pfilter(tt.fn)
			checkSlice(t, tt.expected, actual, tt.name)

			if tt.counter != nil {
				if got := atomic.LoadInt32(tt.counter); got != int32(len(tt.chain)) {
					t.Errorf("TestPfilter %s, counter = %v, want %v", tt.name, got, len(tt.chain))
				}
			}
		})
	}
}
