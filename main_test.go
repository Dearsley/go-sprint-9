package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	var tests = []struct {
		name     string
		size     int
		wantSize int
	}{
		{"zero size", 0, 0},
		{"negative size", -1, 0},
		{"single element", 1, 1},
		{"normal size", 100, 100},
		{"large size", 10000, 10000},
		{"large negative size", -10000, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			elements := generateRandomElements(tt.size)

			require.Len(t, elements, tt.wantSize)

			for i, element := range elements {
				assert.Greater(t, element, 0, "element at index %d", i)
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"nil slice", nil, 0},
		{"single element", []int{42}, 42},
		{"normal slice", []int{3, 8, 2, 10, 5}, 10},
		{"maximum at start", []int{100, 5, 3, 1}, 100},
		{"maximum at end", []int{1, 3, 5, 100}, 100},
		{"equal elements", []int{7, 7, 7, 7}, 7},
		{"all zeros", []int{0, 0, 0}, 0},
		{"negative elements", []int{-10, -2, -30}, -2},
		{"mixed elements", []int{-5, 0, 10, -20}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, maximum(tt.data))
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"nil slice", nil, 0},
		{"single element", []int{42}, 42},
		{"less than 8 elements", []int{3, 7, 2}, 7},
		{"exactly 8 elements", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"9 elements", []int{1, 27, 53, 784000, 455, 26, 547, 18, 100}, 784000},
		{"16 elements", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 16},
		{"17 elements", []int{1000, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 1000},
		{"equal elements", []int{7, 7, 7, 7}, 7},
		{"all zeros", []int{0, 0, 0}, 0},
		{"negative elements", []int{-10, -2, -30}, -2},
		{"mixed elements", []int{-5, 0, 10, -20}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, maxChunks(tt.data))
		})
	}
}
