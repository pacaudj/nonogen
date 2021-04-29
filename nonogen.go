package nonogen

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func NonoGen(sizeY, brightness int, name string) ([][]int, error) {
	img, err := os.Open(name)
	if err != nil {
		fmt.Println("Error while opening file, file doesn't exist.")
		return nil, err
	}
	originImg, err := jpeg.Decode(img)
	if err != nil {
		fmt.Println("Error while opening file, wrong format (jpg/jpeg only).")
		return nil, err
	}
	originImg, err = resizeImg(originImg, sizeY)
	if err != nil {
		return nil, err
	}
	BWImg := imgToGrayScale(originImg)
	bounds := BWImg.Bounds()
	sizeX := bounds.Max.X / (bounds.Max.Y / sizeY)
	sqrSize := bounds.Max.Y / sizeY
	tab := make([][]int, sizeY)
	for i := range tab {
		tab[i] = make([]int, int(sizeX))
	}
	for y := 0; y < bounds.Max.Y; y += sqrSize {
		for x := 0; x < bounds.Max.X; x += sqrSize {
			if avg := getSqrAvgColor(x, y, sqrSize, *BWImg); avg < brightness {
				tab[(y / sqrSize)][(x / sqrSize)] = 1
			}
		}
	}
	return tab, nil
}

func getSqrAvgColor(x, y, size int, img image.RGBA) int {
	avgColor := 0.0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			pixel := img.At(x+j, y+i)
			r, g, b, _ := pixel.RGBA()
			avgColor += 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
		}
	}
	avgColor = avgColor / float64(size) / float64(size)
	return int(avgColor)
}

func SerializeNono(tab [][]int) ([2][][]int, error) {
	row := make([][]int, len(tab))
	column := make([][]int, len(tab[0]))
	for i := range tab {
		nbSquare := 0
		for j := range tab[i] {
			if tab[i][j] == 1 {
				nbSquare++
			} else if nbSquare != 0 {
				row[i] = append(row[i], nbSquare)
				nbSquare = 0
			}
		}
		if nbSquare != 0 {
			row[i] = append(row[i], nbSquare)
		}
	}
	for i := range tab[0] {
		nbSquare := 0
		for j := range tab {
			if tab[j][i] == 1 {
				nbSquare++
			} else if nbSquare != 0 {
				column[i] = append(column[i], nbSquare)
				nbSquare = 0
			}
		}
		if nbSquare != 0 {
			column[i] = append(column[i], nbSquare)
		}
	}
	res := [2][][]int{row, column}
	return res, nil
}
