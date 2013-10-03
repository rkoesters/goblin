package main

import (
	"fmt"
	"os"
)

func main() {
	for _, i := range os.Environ() {
		fmt.Println(i)
	}
}
