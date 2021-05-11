package main

import "fmt"

func FizzBuzz(number int) {
	for i := 1; i <= number; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	FizzBuzz(15)
}
