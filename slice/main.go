package main

import (
	"fmt"
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

// filterSlice filters a slice based on a given predicate and returns a new slice with the matching elements.
func filterSlice(s []int, predicate func(int) bool) []int {
	var result []int
	for _, value := range s {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}

// benchmarkSlice performs the benchmarking of filtering a slice.
func benchmarkSlice(n int) {
	numbers := generateSlice(n)
	evenNumbers := filterSlice(numbers, func(i int) bool { return i%2 == 0 })
	for _, value := range evenNumbers {
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
		benchmarkSlice(NumOfElements)
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
