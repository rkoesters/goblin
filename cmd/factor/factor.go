package main

import (
	"flag"
	"fmt"
	"github.com/rkoesters/goblin/lib/flagutil"
	_ "github.com/rkoesters/goblin/lib/logutil"
	"log"
	"math"
	"strconv"
)

func main() {
	flagutil.Usage = "[number...]"
	flag.Parse()

	for _, i := range flag.Args() {
		n, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}

		factors := factor(n)
		fmt.Println(n)
		for _, j := range factors {
			fmt.Printf("     %v\n", j)
		}
		fmt.Println("")
	}
}

func factor(n int) []int {
	factors := make([]int, 0)
	k := 2

	for n > 1 {
		if n%k == 0 {
			n = n / k
			factors = append(factors, k)
		} else {
			for k++; !isPrime(k); k++ {
			}
		}
	}
	return factors
}

func isPrime(num int) bool {
	for i := 2; float64(i) <= math.Sqrt(float64(num)); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
