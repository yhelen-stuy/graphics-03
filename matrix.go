package main

import (
	"bytes"
	"errors"
	"fmt"
)

type Matrix struct {
	mat        [][]float64
	rows, cols int
	lastcol    int
}

func MakeMatrix(rows, cols int) *Matrix {
	mat := make([][]float64, rows)
	for i := range mat {
		mat[i] = make([]float64, cols)
	}
	matrix := &Matrix{
		mat:  mat,
		rows: rows,
		cols: cols,
	}
	return matrix
}

func (matrix Matrix) Ident() error {
	if matrix.rows != matrix.cols {
		return errors.New("Error: not a square mat")
	}
	for i := range matrix.mat {
		for j := range matrix.mat[i] {
			if i == j {
				matrix.mat[i][j] = 1.0
			} else {
				matrix.mat[i][j] = 0.0
			}
		}
	}
	return nil
}

// Method on a matrix, that modifies the m2, the receiver
// DOES NOT modify m1
// Performs function:
// m1 * m2 -> m2
func (m2 Matrix) Mult(m1 Matrix) error {
	if m2.rows != m1.cols {
		return errors.New("Error: dimensions incompatible")
	}
	for i := 0; i < m1.cols; i++ {
		continue
	}
	return nil
}

func (matrix Matrix) String() string {
	var buf bytes.Buffer
	for i := range matrix.mat {
		for _, c := range matrix.mat[i] {
			buf.WriteString(fmt.Sprintf("%.2f ", c))
		}
		buf.WriteString("\n")
	}
	return buf.String()
}
