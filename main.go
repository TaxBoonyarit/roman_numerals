package main

import (
	"errors"
	"fmt"
)

var NumberRoman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func main() {
	fmt.Println("Hello Roman")
}

func romanNumerals(s string) (int, error) {
	result := 0
	prevValue := 1000 // default max value in roman is `1000`

	// validate count character min and max
	characters := []rune(s)
	if len(characters) <= 0 || len(characters) > 15 {
		return 0, errors.New("invalid value character min : 1 and max : 15")
	}

	// check value roman
	for _, val := range s {
		roman, ok := NumberRoman[string(val)]
		if !ok {
			return 0, errors.New(fmt.Sprintf(`invalid value : %s is not roman number`, string(val)))
		}

		if prevValue < roman {
			result -= prevValue
			result = result + roman - prevValue
		} else {
			result += roman
		}

		prevValue = roman
	}

	return result, nil
}
