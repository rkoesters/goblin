// Print current working directory.
package main

import (
	"fmt"
	_ "github.com/rkoesters/goblin/lib/logutil"
	"log"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pwd)
}
