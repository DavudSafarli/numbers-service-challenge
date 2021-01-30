package main

import (
	"context"
	"strconv"
	"testing"
)

func TestMergerSorter(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected []int
	}{
		{
			input:    [][]int{{1, 2, 3}, {1, 5, 6}, {2, 9, 5}},
			expected: []int{1, 2, 3, 5, 6, 9},
		},
		{
			input:    [][]int{{1, 5, 6}, {-99, 2, 3, 5, 5, 85, 5, 5}, {2, 9, 5}},
			expected: []int{-99, 1, 2, 3, 5, 6, 9, 85},
		},
	}

	for i, tt := range tests {
		t.Run(`test MergerSorter #`+strconv.Itoa(i), func(t *testing.T) {
			subject := NewDefaultMergerSorter()
			testMergerSorter(t, subject, tt.input, tt.expected)
		})
	}
}

func testMergerSorter(t *testing.T, subject MergerSorter, input [][]int, expected []int) {
	actual := subject.MergeAndSort(context.Background(), input)
	if !isEq(actual, expected) {
		t.Error("fail for input", input, "\n",
			"actual   ", actual, "\n",
			"expected ", expected, "\n",
		)
		t.FailNow()
	}
}

func isEq(a, b []int) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
