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

	Q, R := QRFactorization(matrix)
	assert.NotNil(t, Q)
	assert.NotNil(t, R)
	assert.Equal(t, len(Q), len(matrix))
	assert.Equal(t, len(R), len(matrix))
}
