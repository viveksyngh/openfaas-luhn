package function

import (
	"fmt"
	"strconv"
	"strings"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf(strconv.FormatBool(Valid(string(req))))
}

//Extra comment line

//Valid checks if a string given follows luhn's algorithm or not.
func Valid(number string) bool {

	if len(number) <= 1 {
		return false
	} else if len(number) == 2 && strings.HasPrefix(number, " ") {
		return false
	}

	number = strings.Replace(number, " ", "", -1)
	digits := []int{}

	for _, char := range number {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}
		digits = append(digits, digit)
	}

	for i := len(digits) - 2; i >= 0; i = i - 2 {
		if digits[i]*2 > 9 {
			digits[i] = digits[i]*2 - 9
		} else {
			digits[i] = digits[i] * 2
		}
	}

	sum := 0
	for _, digit := range digits {
		sum += digit
	}

	if sum%10 == 0 {
		return true
	}
	return false
}
