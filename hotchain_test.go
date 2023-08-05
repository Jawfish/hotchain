package hotchain

import (
	"testing"
	"time"
)

const (
	chainSize        = 1e5
	fakeWorkDuration = time.Microsecond
)

func setupBenchmarkChain() Chain[int] {
	chain := make(Chain[int], chainSize)
	for i := range chain {
		chain[i] = i
	}
	return chain
}

func square(x int) int { return x * x }

func even(x int) bool { return x%2 == 0 }

func add(x, y int) int { return x + y }

func check(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func checkSlice(t *testing.T, expected, actual Chain[int], msg ...string) {
	if len(expected) != len(actual) {
		t.Errorf("Expected length %v, got length %v: %s", len(expected), len(actual), msg[0])
		return
	}
	for i, val := range expected {
		if val != actual[i] {
			t.Errorf("At index %d: expected %v, got %v", i, val, actual[i])
		}
	}
}

func TestEmptySlice(t *testing.T) {
	tests := []struct {
		name  string
		chain Chain[int]
	}{
		{
			name:  "Operations on empty chain",
			chain: Chain[int]{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check(t, 0, len(tt.chain.Map(square)))
			check(t, 0, len(tt.chain.Filter(even)))
			check(t, 0, tt.chain.Reduce(add, 0))
			check(t, true, tt.chain.All(even))
			check(t, false, tt.chain.Any(even))
		})
	}
}

func TestSingleElementSlice(t *testing.T) {
	tests := []struct {
		name  string
		chain Chain[int]
	}{
		{
			name:  "Operations on single-element chain",
			chain: Chain[int]{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkSlice(t, Chain[int]{4}, tt.chain.Map(square), tt.name+" - Map")
			checkSlice(t, Chain[int]{2}, tt.chain.Filter(even), tt.name+" - Filter")
			check(t, 2, tt.chain.Reduce(add, 0))
			check(t, true, tt.chain.All(even))
			check(t, true, tt.chain.Any(even))
		})
	}
}
