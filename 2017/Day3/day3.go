package main

import (
	"fmt"
)

func main() {

	for c := -5; c < 5; c++ {
		for b := -5; b < 5; b++ {
			fmt.Printf("b: %d; c: %d\n", b, c)
			for i := 0; i < 10; i++ {
				res := (i * i) + b*i + c

				fmt.Printf("Result for i=%d: %d\n", i, res)
			}
		}
	}
}
