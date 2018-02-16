package main

import (
	"fmt"
)

func main() {
	var mat *Matrix
	var id *Matrix
	mat = MakeMatrix(4, 4)
	id = MakeMatrix(4, 4)
	for i := range mat.mat {
		for j := range mat.mat[i] {
			mat.mat[i][j] = 2
		}
	}
	for i := range id.mat {
		for j := range id.mat[i] {
			id.mat[i][j] = float64(i)
		}
	}
	fmt.Println(mat)
	fmt.Println(id)
	mat, _ = mat.Mult(*id)
	fmt.Println(mat)

	// fmt.Println("adding 1,2,3")
	// mat.Ident()
	// mat.AddPoint(1.0, 2.0, 3.0)
	// fmt.Println(mat)

}
