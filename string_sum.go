package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	input = strings.TrimSpace(input)

	if len(input) == 0 {
		return "", fmt.Errorf("nothing to parse: %w", errorEmptyInput)
	}

	sum := 0
	operandsCount := 0
	k := 1

	for i := 0; i < len(input); {
		if input[i] == '-' {
			k *= -1
			i++
		} else if input[i] == '+' || input[i] == ' ' {
			i++
		} else {
			// check if current symbol is a digit
			_, err := strconv.Atoi(input[i : i+1])
			if err != nil {
				return "", fmt.Errorf("invalid input expression: %w", err)
			}

			// find the end of the number
			j := i
			for j < len(input) && err == nil {
				j++
				_, err = strconv.Atoi(input[j:j])
			}

			// convert string number to int
			n, err := strconv.Atoi(input[i:j])
			if err != nil {
				return "", fmt.Errorf("could not convert to int: %w", err)
			}

			sum += k * n
			operandsCount++
			k = 1

			i = j
		}
	}

	if operandsCount != 2 {
		return "", fmt.Errorf("invalid number of operands: %w", errorNotTwoOperands)
	}

	output = strconv.Itoa(sum)
	return output, nil
}
