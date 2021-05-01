package nonogen_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/pacaudj/nonogen"
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
		_, err := nonogen.NonoGen(table.size, table.bright, table.name)
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
		{10, 30000, "test-pic/bird.jpg", "bird_10.jpg"},
		{15, 30000, "test-pic/bird.jpg", "bird_15.jpg"},
		{20, 30000, "test-pic/bird.jpg", "bird_20.jpg"},
		{30, 25000, "test-pic/shosta.jpg", "shosta_1.jpg"},
		{30, 30000, "test-pic/shosta.jpg", "shosta_2.jpg"},
		{30, 35000, "test-pic/shosta.jpg", "shosta_3.jpg"},
		{20, 30000, "test-pic/lion.jpg", "lion_1.jpg"},
		{30, 30000, "test-pic/lion.jpg", "lion_2.jpg"},
		{55, 30000, "test-pic/lion.jpg", "lion_3.jpg"},
	}
	for _, table := range tables {
		_, err := nonogen.NonoGen(table.size, table.bright, table.name)
		if err != nil {
			os.Exit(1)
		} else {
			fmt.Println("Successfully generated nonogram", table.dest)
		}
	}
}
