package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	d, err := os.Getwd()
	if err != nil {
		fmt.Print("???")
	} else {
		fmt.Print(filepath.Base(d))
	}
}
