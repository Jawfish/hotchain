package hotchain

import "testing"

func TestEach(t *testing.T) {
	tests := []struct {
		name     string
		chain    Chain[int]
		fn       func(*int, int)
		expected int
	}{
		{
			name:     "Sum all numbers",
			chain:    Chain[int]{1, 2, 3, 4, 5},
			fn:       func(sum *int, x int) { *sum += x },
			expected: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum := 0
			tt.chain.Each(func(x int) { tt.fn(&sum, x) })
			check(t, tt.expected, sum)
		})
	}
}
