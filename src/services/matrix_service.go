package services

import (
	"gonum.org/v1/gonum/mat"
)

// Funcion que factoriza en QR usando la Libreria gonun
func QRFactorization(matrix [][]float64) ([][]float64, [][]float64) {
	rows := len(matrix)
	cols := len(matrix[0])
	flat := make([]float64, 0, rows*cols)
	for i := range matrix {
		flat = append(flat, matrix[i]...)
	}
	m := mat.NewDense(rows, cols, flat)

	var qr mat.QR
	qr.Factorize(m)

	Q := mat.NewDense(rows, cols, nil)
	R := mat.NewDense(cols, cols, nil)
	qr.QTo(Q)
	qr.RTo(R)

	q := make([][]float64, rows)
	r := make([][]float64, cols)
	for i := 0; i < rows; i++ {
		q[i] = Q.RawRowView(i)
	}
	for i := 0; i < cols; i++ {
		r[i] = R.RawRowView(i)
	}

	return q, r
}

// Funcion que rota la matriz que recibe
func RotateMatrix(matrix [][]float64) [][]float64 {
	if len(matrix) == 0 {
		return matrix
	}
	rowCount := len(matrix)
	colCount := len(matrix[0])
	rotated := make([][]float64, colCount)
	for i := range rotated {
		rotated[i] = make([]float64, rowCount)
	}
	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			rotated[j][rowCount-1-i] = matrix[i][j]
		}
	}
	return rotated
}
