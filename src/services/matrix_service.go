package services

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func QRFactorization(matrix [][]float64) ([][]float64, [][]float64, error) {
	rows := len(matrix)
	if rows == 0 {
		return nil, nil, fmt.Errorf("la matriz no puede estar vacía")
	}
	cols := len(matrix[0])
	if cols == 0 {
		return nil, nil, fmt.Errorf("la matriz no puede tener columnas vacías")
	}
	for _, row := range matrix {
		if len(row) != cols {
			return nil, nil, fmt.Errorf("todas las filas deben tener el mismo número de columnas")
		}
	}

	if rows < cols {
		return nil, nil, fmt.Errorf("la matriz debe tener al menos tantas filas como columnas para la factorización QR")
	}

	flat := make([]float64, 0, rows*cols)
	for i := range matrix {
		flat = append(flat, matrix[i]...)
	}

	m := mat.NewDense(rows, cols, flat)

	var qr mat.QR
	qr.Factorize(m)

	Q := mat.NewDense(rows, rows, nil)
	R := mat.NewDense(rows, cols, nil)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error al copiar las matrices Q y R:", r)
		}
	}()

	qr.QTo(Q)
	qr.RTo(R)

	qRes := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		qRes[i] = make([]float64, rows)
		copy(qRes[i], Q.RawRowView(i))
	}

	rRes := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		rRes[i] = make([]float64, cols)
		copy(rRes[i], R.RawRowView(i))
	}

	return qRes, rRes, nil
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
