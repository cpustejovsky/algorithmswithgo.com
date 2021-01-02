package module01

import (
	"fmt"
)

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {
	for i := 1; i < n; i++ {
		fmt.Printf("%v, ", determineFizzBuzz(i))
	}
	fmt.Println(determineFizzBuzz(n))
}

func determineFizzBuzz(i int) string {
	switch {
	case i%15 == 0:
		return "Fizz Buzz"
	case i%5 == 0:
		return "Buzz"
	case i%3 == 0:
		return "Fizz"
	default:
		return fmt.Sprint(i)
	}
}
