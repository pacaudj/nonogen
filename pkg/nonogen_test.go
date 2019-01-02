package pkg

import (
	"fmt"
	"github.com/shadonovitch/nonogen/pkg/draw"
	"os"
	"testing"
)

func TestErrorCaseNonogen(b *testing.T) {
	tables := []struct {
		size   int
		bright int
		name   string
		test   string
	}{
		{25, 30000, "test-pic/shosta", "File doesn't exist"},
		{25, 30000, "test-pic/shosta.png", "Wrong file format"},
		{250, 30000, "test-pic/shosta.jpg", "Size is too big"},
		{2, 30000, "test-pic/shosta.jpg", "Size is too small"},
	}
	for _, table := range tables {
		_, err := Nonogen(table.size, table.bright, table.name)
		if err == nil {
			fmt.Println(table.test, ": FAILED")
			os.Exit(1)
		} else {
			fmt.Println(table.test, ": SUCCESS")
		}
	}
}

func TestNonogen(b *testing.T) {
	tables := []struct {
		size   int
		bright int
		name   string
		dest   string
	}{
		{10, 30000, "test-pic/bird.jpg", "test-output/bird_10.jpg"},
		{15, 30000, "test-pic/bird.jpg", "test-output/bird_15.jpg"},
		{20, 30000, "test-pic/bird.jpg", "test-output/bird_20.jpg"},
		{30, 25000, "test-pic/shosta.jpg", "test-output/shosta_1.jpg"},
		{30, 30000, "test-pic/shosta.jpg", "test-output/shosta_2.jpg"},
		{30, 35000, "test-pic/shosta.jpg", "test-output/shosta_3.jpg"},
		{20, 30000, "test-pic/lion.jpg", "test-output/lion_1.jpg"},
		{30, 30000, "test-pic/lion.jpg", "test-output/lion_2.jpg"},
		{55, 30000, "test-pic/lion.jpg", "test-output/lion_3.jpg"},
	}
	for _, table := range tables {
		nono, err := Nonogen(table.size, table.bright, table.name)
		if err != nil {
			os.Exit(1)
		} else {
			draw.NonoToJPG(nono, table.dest)
			fmt.Println("Successfully generated file", table.dest)
		}
	}
}