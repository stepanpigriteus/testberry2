package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	s := "я3ыв!4.2   s"
	str, err := unpackString(s)
	fmt.Println(str, err)
}

func unpackString(input string) (string, error) {

	errInvalidString := errors.New("invalid str: contains only digits or incorrect format")
	if input == "" {
		return input, nil
	}

	runes := []rune(input)
	var result strings.Builder
	for i := 0; i < len(runes); i++ {
		r := runes[i]

		if r == '\\' {
			if i+1 >= len(runes) {
				return "", errInvalidString
			}
			char := runes[i+1]
			i++
			count := 1
			if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
				digitStr := ""
				for j := i + 1; j < len(runes) && unicode.IsDigit(runes[j]); j++ {
					digitStr += string(runes[j])
					i = j
				}
				num, err := strconv.Atoi(digitStr)
				if err != nil {
					return "", errInvalidString
				}
				count = num
			}
			for c := 0; c < count; c++ {
				result.WriteRune(char)
			}
			continue
		}

		if unicode.IsDigit(r) {
			return "", errInvalidString
		}

		char := r
		count := 1
		if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
			digitStr := ""
			for j := i + 1; j < len(runes) && unicode.IsDigit(runes[j]); j++ {
				digitStr += string(runes[j])
				i = j
			}
			num, err := strconv.Atoi(digitStr)
			if err != nil {
				return "", errInvalidString
			}
			count = num
		}
		for c := 0; c < count; c++ {
			result.WriteRune(char)
		}
	}

	return result.String(), nil
}
