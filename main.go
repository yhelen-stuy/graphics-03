package main

import (
	"fmt"
	// "math/rand"
	// "time"
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
	fmt.Println("Making Matrix full of 2s")
	fmt.Println(mat)
	fmt.Println("Matrix where all values in a row are equal to the row they are in")
	fmt.Println(id)
	mat, _ = mat.Mult(*id)
	fmt.Println("Multiplying the previous two should make a mat of all 12s")
	fmt.Println(mat)

	fmt.Println("Adding points (1,2,3) and (4,5,6) to identity")
	mat.Ident()
	mat.AddPoint(1.0, 2.0, 3.0)
	mat.AddPoint(4.0, 5.0, 6.0)
	fmt.Println(mat)

	fmt.Println("Draw lines from (0,0) to (250,250) to (300,400)")
	image := MakeImage(500, 500)
	edges := MakeMatrix(4, 0)
	edges.AddEdge(0.0, 0.0, 0.0, 250.0, 250.0, 0.0)
	edges.AddEdge(250.0, 250.0, 0.0, 300.0, 400.0, 0.0)
	fmt.Println(edges)
	image.DrawLines(edges, Color{r: 255, b: 0, g: 0})
	image.SavePPM("mat.ppm")
	/* My gallery code
	rand.Seed(int64(time.Now().Unix()))
	mat := MakeMatrix(4, 0)
	image := MakeImage(500, 500)
	for i := 0; i < 20; i++ {
		mat.AddEdge(rand.Float64()*500, 500-rand.Float64()*500, 0.0, rand.Float64()*500, 500-rand.Float64()*500, 0.0)
	}
	image.DrawLines(mat, Color{r: 0, b: 0, g: 0})
	image.SavePPM("mat.ppm")
	*/
}
