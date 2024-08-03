package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateMatrix(t *testing.T) {
	tests := []struct {
		input    [][]float64
		expected [][]float64
	}{
		{
			input: [][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]float64{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			input: [][]float64{
				{1, 2},
				{3, 4},
				{5, 6},
			},
			expected: [][]float64{
				{5, 3, 1},
				{6, 4, 2},
			},
		},
		{
			input: [][]float64{},
			expected: [][]float64{},
		},
		{
			input: [][]float64{
				{1},
			},
			expected: [][]float64{
				{1},
			},
		},
	}

	for _, test := range tests {
		result := RotateMatrix(test.input)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}
