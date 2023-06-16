package main

import (
	"fmt"
)

func main() {
	fizzBuzz()
}

func fizzBuzz() {

	for n := 1; n <= 20; n++ {
		switch {
		case n%3 == 0 && n%5 == 0:
			fmt.Println("fizz buzz")
			// break
		case n%3 == 0:
			fmt.Println("fizz ")
		case n%5 == 0:
			fmt.Println("buzz ")
		default:
			fmt.Println(n)
		}
	}
}
