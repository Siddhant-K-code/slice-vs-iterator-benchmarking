package main

import (
	"fmt"
	"iter"
	"math"
	"runtime"
)

const (
	NumOfElements = 10_000_000
)

// generateSlice creates a slice of integers from 0 to n-1.
func generateSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = i
	}
	return s
}

// filterIterator returns an iterator that filters a slice based on a given predicate.
func filterIterator(seq []int, predicate func(int) bool) iter.Seq[int] {
	return func(yield func(int) bool) {
		for _, value := range seq {
			if predicate(value) {
				if !yield(value) {
					return
				}
			}
		}
	}
}

// benchmarkIterator performs the benchmarking of filtering a slice using an iterator.
func benchmarkIterator(n int) {
	numbers := generateSlice(n)
	evenNumbers := filterIterator(numbers, func(i int) bool { return i%2 == 0 })
	for value := range evenNumbers {
		_ = value
	}
}

// max returns the maximum of two uint64 numbers.
func max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

// min returns the minimum of two uint64 numbers.
func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	var averageCount uint64
	var averageTotalAlloc uint64
	var maxTotalAlloc uint64 = 0
	var minTotalAlloc uint64 = math.MaxUint64
	var averageMallocs uint64
	var maxMallocs uint64 = 0
	var minMallocs uint64 = math.MaxUint64

	var m1, m2 runtime.MemStats
	for i := 1; i <= 10_000; i++ {
		if i%1000 == 0 {
			fmt.Printf("Running iteration: %d\n", i)
		}
		runtime.GC()
		runtime.ReadMemStats(&m1)
		benchmarkIterator(NumOfElements)
		runtime.ReadMemStats(&m2)

		diffTotalAlloc := m2.TotalAlloc - m1.TotalAlloc
		diffMallocs := m2.Mallocs - m1.Mallocs

		averageTotalAlloc = (averageTotalAlloc*averageCount + diffTotalAlloc) / uint64(i)
		maxTotalAlloc = max(maxTotalAlloc, diffTotalAlloc)
		minTotalAlloc = min(minTotalAlloc, diffTotalAlloc)

		averageMallocs = (averageMallocs*averageCount + diffMallocs) / uint64(i)
		maxMallocs = max(maxMallocs, diffMallocs)
		minMallocs = min(minMallocs, diffMallocs)

		averageCount++
	}

	fmt.Printf("NumOfElements: %d\n", NumOfElements)
	fmt.Println()
	fmt.Printf("AverageTotalAlloc: %d B/op\n", averageTotalAlloc)
	fmt.Printf("MaxTotalAlloc: %d B\n", maxTotalAlloc)
	fmt.Printf("MinTotalAlloc: %d B\n", minTotalAlloc)
	fmt.Println()
	fmt.Printf("AverageMallocs: %d allocs/op\n", averageMallocs)
	fmt.Printf("MaxMallocs: %d allocs\n", maxMallocs)
	fmt.Printf("MinMallocs: %d allocs\n", minMallocs)
}
