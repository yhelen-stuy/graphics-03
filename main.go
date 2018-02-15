package main

import (
	"fmt"
)

func main() {
	var mat *Matrix
	var id *Matrix
	mat = MakeMatrix(2, 2)
	id = MakeMatrix(2, 2)
	id.Ident()
	fmt.Println(id)
	for i := range mat.mat {
		for j := range mat.mat[i] {
			mat.mat[i][j] = 2
		}
	}
	fmt.Println(mat)
	mat.Mult(*id)
	fmt.Println(mat)
}
