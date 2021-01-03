package module01

import "fmt"

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//

type CharSet map[string]int

func CreateCharSet(chars string) CharSet {
	c := make(CharSet)
	for i, v := range chars {
		c[string(v)] = i
	}
	return c
}

var charset = CreateCharSet("0123456789ABCDEF")

func BaseToDec(value string, base int) int {
	// res := 0
	// multiplier := 1
	// for i := len(value) - 1; i >= 0; i-- {
	// 	index := charset[string(rune(value[i]))]
	// 	res += index * multiplier
	// 	multiplier *= base
	// }
	// return res
	if len(value) == 0 {
		return 0
	}
	num := charset[string(rune(value[0]))]
	fmt.Println("num from charset: ", num)
	power := len(value) - 1
	fmt.Println("base: ", base)
	fmt.Println("power: ", power)
	result := num * (base ^ power)
	fmt.Println("result: ", result)
	return result + BaseToDec(value[1:], base)
}
