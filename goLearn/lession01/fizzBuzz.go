package main

import (
	"fmt"
)

func main() {
	// Challenge 01
	fizzBuzz()

	// Challenge 02
	evenEnded()

}

func evenEnded() {
	fmt.Println("================== Begin Challenge 02 ===================")

	count := 0

	for x := 1000; x <= 9999; x++ {
		for y := x; y <= 9999; y++ {
			num := x * y
			numString := fmt.Sprintf("%d", num)
			if numString[0] == numString[len(numString)-1] {
				count++
			}
		}
	}

	fmt.Println(count)
	fmt.Println("================== End Challenge 02 ===================")
}

func fizzBuzz() {

	fmt.Println("================== Begin Challenge 01 ===================")
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
	fmt.Println("================== End Challenge 01 ===================")
}
