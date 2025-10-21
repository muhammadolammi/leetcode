package main

import (
	"fmt"
	"strings"
)

func rcheck(i, j int, s, p []byte, dp map[string]bool) bool {
	key := fmt.Sprintf("%d,%d", i, j)
	if val, ok := dp[key]; ok {
		return val
	}

	if i >= len(s) && j >= len(p) {
		dp[key] = true
		return true
	}
	if j >= len(p) {
		dp[key] = false
		return false
	}

	match := i < len(s) && (s[i] == p[j] || p[j] == '.')

	if j+1 < len(p) && p[j+1] == '*' {
		// * means: skip the pattern or use it if match
		res := rcheck(i, j+2, s, p, dp) ||
			(match && rcheck(i+1, j, s, p, dp))
		dp[key] = res
		return res
	}

	if match {
		res := rcheck(i+1, j+1, s, p, dp)
		dp[key] = res
		return res
	}

	dp[key] = false
	return false
}

func isMatch(s string, p string) bool {

	if p == ".*" {
		return true
	}
	// when p doesnt contain . and * its string matching
	if !strings.Contains(p, ".") && !strings.Contains(p, "*") {
		return s == p
	}
	sBytes := []byte(s)
	pBytes := []byte(p)

	dp := map[string]bool{}
	match := rcheck(0, 0, sBytes, pBytes, dp)
	return match

}
