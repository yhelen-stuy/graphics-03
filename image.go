package main

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"os"
)

type Color struct {
	r int
	g int
	b int
}

type Image struct {
	img    [][]Color
	height int
	width  int
}

func MakeImage(height, width int) *Image {
	img := make([][]Color, height)
	for i := range img {
		img[i] = make([]Color, width)
	}
	image := &Image{
		img:    img,
		height: height,
		width:  width,
	}
	image.Clear()
	return image
}

func (image Image) plot(c Color, x, y int) error {
	if x < 0 || x > image.height || y < 0 || y > image.width {
		return errors.New("Error: Coordinate invalid")
	}
	image.img[x][y] = c
	return nil
}

func (image Image) fill(c Color) {
	for y := 0; y < image.width; y++ {
		for x := 0; x < image.height; x++ {
			image.plot(c, x, y)
		}
	}
}

func (image Image) Clear() {
	image.fill(Color{r: 255, g: 255, b: 255})
}

func (image Image) SavePPM(filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	// TODO: Take variant max color
	buffer.WriteString(fmt.Sprintf("P3 %d %d 255\n", image.height, image.width))

	for y := 0; y < image.width; y++ {
		for x := 0; x < image.height; x++ {
			newY := image.width - 1 - y
			buffer.WriteString(fmt.Sprintf("%d %d %d ", image.img[x][newY].r, image.img[x][newY].g, image.img[x][newY].b))
		}
		buffer.WriteString("\n")
	}

	f.WriteString(buffer.String())
	f.Close()
	return nil
}

func (image Image) DrawLine(c Color, x0, y0, x1, y1 int) error {
	if x0 > x1 {
		x1, x0 = x0, x1
		y1, y0 = y0, y1
	}

	// fmt.Printf("Drawing (%d, %d) to (%d, %d)\n", x0, y0, x1, y1)

	deltaX := x1 - x0
	deltaY := y1 - y0
	if deltaY >= 0 {
		if math.Abs(float64(deltaY)) <= math.Abs(float64(deltaX)) {
			// fmt.Println("OctantI")
			image.drawLineOctantI(c, deltaY, deltaX*-1, x0, y0, x1, y1)
		} else {
			// fmt.Println("OctantII")
			image.drawLineOctantII(c, deltaY, deltaX*-1, x0, y0, x1, y1)
		}
	} else {
		if math.Abs(float64(deltaY)) > math.Abs(float64(deltaX)) {
			// fmt.Println("OctantVII")
			image.drawLineOctantVII(c, deltaY, deltaX*-1, x0, y0, x1, y1)
		} else {
			// fmt.Println("OctantVIII")
			image.drawLineOctantVIII(c, deltaY, deltaX*-1, x0, y0, x1, y1)
		}
	}
	return nil
}

func (image Image) drawLineOctantI(c Color, lA, lB, x0, y0, x1, y1 int) error {
	y := y0
	lD := 2*lA + lB
	for x := x0; x < x1; x++ {
		err := image.plot(c, x, y)
		if err != nil {
			return err
		}
		if lD > 0 {
			y++
			lD += 2 * lB
		}
		lD += 2 * lA
	}
	return nil
}

func (image Image) drawLineOctantII(c Color, lA, lB, x0, y0, x1, y1 int) error {
	x := x0
	lD := lA + 2*lB
	for y := y0; y < y1; y++ {
		err := image.plot(c, x, y)
		if err != nil {
			return err
		}
		if lD < 0 {
			x++
			lD += 2 * lA
		}
		lD += 2 * lB
	}
	return nil
}

func (image Image) drawLineOctantVII(c Color, lA, lB, x0, y0, x1, y1 int) error {
	x := x0
	lD := lA - 2*lB
	for y := y0; y > y1; y-- {
		// fmt.Printf("(%d, %d)\n", x, y)
		err := image.plot(c, x, y)
		if err != nil {
			return err
		}
		if lD > 0 {
			x++
			lD += 2 * lA
		}
		lD -= 2 * lB
	}
	return nil
}

func (image Image) drawLineOctantVIII(c Color, lA, lB, x0, y0, x1, y1 int) error {
	y := y0
	lD := 2*lA - lB
	for x := x0; x < x1; x++ {
		// fmt.Printf("(%d, %d)\n", x, y)
		err := image.plot(c, x, y)
		if err != nil {
			return err
		}
		if lD < 0 {
			y--
			lD -= 2 * lB
		}
		lD += 2 * lA
	}
	return nil
}
