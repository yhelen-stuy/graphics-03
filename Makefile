all:
	go run image.go main.go matrix.go

run:
	display lines.ppm

clean:
	rm *.ppm
