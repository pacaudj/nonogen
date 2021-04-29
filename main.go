package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pacaud_j/nonogen"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: nonogen size brightness source (dest)")
		os.Exit(1)
	}
	size, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Usage: nonogen size brightness source (dest)")
		return
	}
	brightness, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Usage: nonogen size brightness source (dest)")
		return
	}
	fileName := os.Args[3]
	nono, err := nonogen.NonoGen(size, brightness, fileName)
	if err != nil {
		return
	}
	if len(os.Args) >= 5 {
		destName := os.Args[4]
		nonogen.NonoToJPG(nono, destName)
	}
	nonogen.NonoToJPG(nono, "test_nono.jpg")
	serializedNono, err := nonogen.SerializeNono(nono)
	nonogen.IsSolvable(serializedNono)
}
