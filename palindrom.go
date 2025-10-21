package main

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	xStrings := strconv.Itoa(x)
	reversedXStrings := string(reverseString([]byte(xStrings)))
	return xStrings == reversedXStrings
}
