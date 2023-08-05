package hotchain

import (
	"testing"
	"time"
)

func BenchmarkMap(b *testing.B) {
	chain := setupBenchmarkChain()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		chain.Map(func(x int) int {
			return x * x
		})
	}
}

func BenchmarkFilter(b *testing.B) {
	chain := setupBenchmarkChain()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		chain.Filter(func(x int) bool {
			return x%2 == 0
		})
	}
}
func BenchmarkPmap(b *testing.B) {
	chain := setupBenchmarkChain()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		chain.Pmap(func(x int) int {
			return x * x
		})
	}
}

func BenchmarkPfilter(b *testing.B) {
	chain := setupBenchmarkChain()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		chain.Pfilter(func(x int) bool {
			return x%2 == 0
		})
	}
}

func BenchmarkMapSimulatedWork(b *testing.B) {
	chain := setupBenchmarkChain()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		chain.Map(func(x int) int {
			time.Sleep(fakeWorkDuration)
			return x * x
		})
	}
}

func BenchmarkFilterSimulatedWork(b *testing.B) {
	chain := setupBenchmarkChain()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		chain.Filter(func(x int) bool {
			time.Sleep(fakeWorkDuration)
			return x%2 == 0
		})
	}
}

func BenchmarkPmapSimulatedWork(b *testing.B) {
	chain := setupBenchmarkChain()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		chain.Pmap(func(x int) int {
			time.Sleep(fakeWorkDuration)
			return x * x
		})
	}
}

func BenchmarkPfilterSimulatedWork(b *testing.B) {
	chain := setupBenchmarkChain()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		chain.Pfilter(func(x int) bool {
			time.Sleep(fakeWorkDuration)
			return x%2 == 0
		})
	}
}
