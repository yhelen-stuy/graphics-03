package main

import ()

type Matrix struct {
	mat        [][]float64
	rows, cols int
	lastcol    int
}

func MakeMatrix(rows, cols int) *Matrix {
	mat = make([][]float64, rows)
	for i := range mat {
		mat[i] = make([]float64, cols)
	}
	matrix = &Matrix{
		mat:  mat,
		rows: rows,
		cols: cols,
	}
	return matrix
}
