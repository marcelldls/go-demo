package main

import "strconv"

// Check if number divisble by 2
func isEven(input int) string {
	input_str := strconv.Itoa(input) // Int to ascii
	isEven := (input % 2) == 0
	if isEven {
		return input_str + " is even"
	}
	return input_str + " is odd"
}
