package finder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTheMissingLength(t *testing.T) {
	tests := []struct {
		given    [][]int
		expected int
	}{
		{
			given:    [][]int{[]int{1}, []int{1, 2}, []int{1, 2, 3, 4}},
			expected: 3,
		},
		{
			given:    [][]int{[]int{1, 2}, []int{1, 2, 3}, []int{1, 2, 3, 4, 5}},
			expected: 4,
		},
		{
			given:    [][]int{[]int{1, 2}, []int{1, 2, 3, 4}},
			expected: 3,
		},
		{
			given:    [][]int{[]int{1, 2, 3}, []int{1}, []int{1, 2, 3, 4}},
			expected: 2,
		},
		{
			given:    [][]int{}, // empty array => return 0
			expected: 0,
		},
		{
			given:    nil, // nil array => return 0
			expected: 0,
		},
		{
			given:    [][]int{[]int{1, 2, 3}, []int{1}, []int{1, 2, 3, 4}},
			expected: 2,
		},
		{
			given:    [][]int{[]int{1, 2, 3}, []int{1, 2, 3, 4}, nil}, // nil element => return 0
			expected: 0,
		},
		{
			given:    [][]int{[]int{1, 2, 3}, []int{}, []int{1, 2, 3, 4}}, // empty element => return 0
			expected: 0,
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expected, FindTheMissingLength(tc.given))
	}
}

func BenchmarkFindTheMissingLength(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		FindTheMissingLength([][]int{})
	}
}
