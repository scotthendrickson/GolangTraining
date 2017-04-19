package main

import (
	"fmt"
)

func main() {
	x := 13 % 3
	fmt.Println(x)

	for i := 0; i < 100; i++ {
		y := i % 2
		if y == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}