package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQRFactorization(t *testing.T) {
	matrix := [][]float64{
		{1, 2},
		{3, 4},
	}

	Q, R, err := QRFactorization(matrix)
	assert.NoError(t, err)
	assert.NotNil(t, Q)
	assert.NotNil(t, R)
	assert.Equal(t, len(Q), len(matrix))
	assert.Equal(t, len(Q[0]), len(matrix))
	assert.Equal(t, len(R), len(matrix))
	assert.Equal(t, len(R[0]), len(matrix[0]))
}
func TestQRFactorization_DifferentRowLengths(t *testing.T) {
	matrix := [][]float64{
		{1, 2},
		{3, 4, 5},
	}

	Q, R, err := QRFactorization(matrix)
	assert.Error(t, err)
	assert.Nil(t, Q)
	assert.Nil(t, R)
}
func TestQRFactorization_EmptyColumns(t *testing.T) {
	matrix := [][]float64{
		{},
		{},
	}

	Q, R, err := QRFactorization(matrix)
	assert.Error(t, err)
	assert.Nil(t, Q)
	assert.Nil(t, R)
}
