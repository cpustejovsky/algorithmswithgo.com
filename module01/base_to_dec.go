package module01

import (
	"math"
)

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
	if len(value) == 0 {
		return 0
	}
	num := charset[string(rune(value[0]))]
	power := len(value) - 1
	result := float64(num) * math.Pow(float64(base), float64(power))
	return int(result) + BaseToDec(value[1:], base)
}
