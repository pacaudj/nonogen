package nonogen

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

func drawLine(x1, x2, y int, img image.RGBA, color color.Color) {
	for ; x1 < x2; x1++ {
		img.Set(x1, y, color)
	}
}

func drawVerticalLine(y1, y2, x int, img image.RGBA, color color.Color) {
	for ; y1 < y2; y1++ {
		img.Set(x, y1, color)
	}
}

func drawRect(x, y, x2, y2 int, img image.RGBA, color color.Color) {
	for ; y < y2; y++ {
		drawLine(x, x2, y, img, color)
	}
}

func imgToGrayScale(originImg image.Image) *image.RGBA {
	bounds := originImg.Bounds()
	BWImg := image.NewRGBA(bounds)
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			oldPixel := originImg.At(x, y)
			pixel := color.GrayModel.Convert(oldPixel)
			BWImg.Set(x, y, pixel)
		}
	}
	return BWImg
}

func NonoToJPG(nono [][]int, name string) {
	size := 10
	sizeY := len(nono)
	sizeX := len(nono[0])
	img := image.NewRGBA(image.Rect(0, 0, sizeX*size, sizeY*size))
	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			if nono[y][x] == 1 {
				drawRect(x*size, y*size, x*size+size, y*size+size, *img, color.Black)
			} else {
				drawRect(x*size, y*size, x*size+size, y*size+size, *img, color.White)
			}
		}
	}
	for i := 1; i < sizeY; i++ {
		drawLine(0, size*sizeX, i*size, *img, color.RGBA{255, 0, 0, 1})
	}
	for i := 1; i < sizeX; i++ {
		drawVerticalLine(0, size*sizeY, i*size, *img, color.RGBA{255, 0, 0, 1})
	}
	JpgDraw(name, img)
}

func JpgDraw(name string, img image.Image) {
	var opt jpeg.Options
	opt.Quality = 75
	out, err := os.Create(name)
	err = jpeg.Encode(out, img, &opt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func resizeImg(originImg image.Image, size int) (image.Image, error) {
	bounds := originImg.Bounds()
	if size > bounds.Max.Y {
		fmt.Println("Requested size is too big")
		return nil, errors.New("Requested size is too big")
	}
	if size < 5 {
		fmt.Println("Requested size is too small")
		return nil, errors.New("Requested size is too small")
	}
	sizeX := bounds.Max.X / (bounds.Max.Y / size)
	desiredX := (bounds.Max.Y / size) * sizeX
	desiredY := (bounds.Max.Y / size) * size
	originImg = resize.Resize(uint(desiredX), uint(desiredY), originImg, resize.Lanczos3)
	return originImg, nil
}
