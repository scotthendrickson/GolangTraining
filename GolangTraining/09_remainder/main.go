package main

import (
	"fmt"
)

func main() {
	x := 13 % 3
	fmt.Println(x)

	for i := 0; i < 100; i++ {
		y := i % 2
		fizzBuzz(i)
		if y == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}

func fizzBuzz(num int) {
	switch {
	case num%3 == 0:
		fmt.Println("fizz")
	case num%5 == 0:
		fmt.Println("Buzz")
	case num%15 == 0:
		fmt.Println("FizzBuzz")
	default:
		fmt.Println(num)
	}
}