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

func (m Matrix) Ident() error {
	if m.rows != m.cols {
		return errors.New("Error: not a square mat")
	}
	for i := range m.mat {
		for j := range m.mat[i] {
			if i == j {
				m.mat[i][j] = 1.0
			} else {
				m.mat[i][j] = 0.0
			}
		}
	}
	return nil
}

// Method on a matrix, that modifies m2, the receiver
// DOES NOT modify m1
// Performs function:
// m1 * m2 -> m2
func (m2 Matrix) Mult(m1 Matrix) (*Matrix, error) {
	// TODO: Make sure it works on non identity matrix
	if m1.rows != m2.cols {
		return nil, errors.New("Error: dimensions incompatible")
	}
	prod := MakeMatrix(m2.rows, m1.cols)
	value := 0.0
	for r1 := 0; r1 < m1.rows; r1++ {
		for c2 := 0; c2 < m2.cols; c2++ {
			for rc := 0; rc < m1.cols; rc++ {
				value += m2.mat[r1][rc] * m1.mat[rc][c2]
			}
			prod.mat[r1][c2] = value
			value = 0.0
		}
	}
	return prod, nil
}

func (m Matrix) AddPoint(x, y, z float64) {
	p := make([]float64, 4)
	p[0] = x
	p[1] = y
	p[2] = z
	p[3] = 1
	m.mat = append(m.mat, p)
}

func (m Matrix) String() string {
	var buf bytes.Buffer
	for i := range m.mat {
		for _, c := range m.mat[i] {
			buf.WriteString(fmt.Sprintf("%.2f ", c))
		}
		buf.WriteString("\n")
	}
	return buf.String()
}
