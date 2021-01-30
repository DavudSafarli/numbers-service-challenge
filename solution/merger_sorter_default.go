package main

import (
	"context"
	"sort"
)

// DefaultMergerSorter is a default, non-optimized implementation of MergerSorter
type DefaultMergerSorter struct {
}

// NewDefaultMergerSorter is a constructor for DefaultMergerSorter
func NewDefaultMergerSorter() DefaultMergerSorter {
	return DefaultMergerSorter{}
}

// MergeAndSort merges and sortes slice of slices in the simplest way
func (ms DefaultMergerSorter) MergeAndSort(ctx context.Context, slices [][]int) []int {
	ans := make([]int, 0)
	for _, slice := range slices {
		for _, element := range slice {
			if existsInSlice(ans, element) {
				continue
			}
			ans = append(ans, element)
		}
	}
	sort.Ints(ans)
	return ans
}

func existsInSlice(slice []int, needle int) bool {
	for _, element := range slice {
		if element == needle {
			return true
		}
	}

	return false
}
