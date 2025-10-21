package main

import (
	"math"
	"strings"
)

func reverseString(s []byte) []byte {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left += 1
		right -= 1
	}
	return s
}

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	isNegative := false
	if s == "" {
		return 0
	}
	if string(s[0]) == "-" {
		isNegative = true
	}
	digits := map[string]int64{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}
	validSring := ""
	for i, value := range s {
		stringedValue := string(value)
		if i == 0 && (stringedValue == "-" || stringedValue == "+") {
			continue
		}

		if stringedValue == " " {
			break
		}
		_, valid := digits[stringedValue]
		if !valid {
			break
		}
		validSring += stringedValue
	}
	returnValue := float64(0)
	for i, value := range reverseString([]byte(validSring)) {

		stringedValue := string(value)

		if stringedValue == "0" {
			continue
		}
		intValue := digits[stringedValue]
		power := math.Pow10(i)
		returnValue += float64(intValue) * power
	}
	if isNegative {
		returnValue = -returnValue
	}
	if returnValue > math.MaxInt32 {
		return math.MaxInt32
	}
	if returnValue < math.MinInt32 {
		return math.MinInt32
	}
	return int(returnValue)

}
