all:
	go run image.go main.go matrix.go

run:
	display mat.ppm

clean:
	rm *.ppm
