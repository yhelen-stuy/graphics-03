package main

import (
	"fmt"
)

func main() {
	var mat *Matrix
	mat = MakeMatrix(4, 4)
	fmt.Println(mat)
	mat.Ident()
	fmt.Println(mat)
}
