package main

// func rcheck(i, j int, s, p []byte, iterations *int) bool {
// 	*iterations++
// 	// fmt.Println(i, j)

// 	if i >= len(s) && j >= len(p) {
// 		return true
// 	}
// 	if j >= len(p) {
// 		//  we are done with p but s still remain so no match
// 		return false
// 	}
// 	match := false
// 	// lets check if i is still in bound
// 	if i < len(s) {
// 		match = (s[i] == p[j]) || string(p[j]) == "."
// 	}
// 	// fmt.Println("match: ", match)
// 	// lets see if we can go further in the p
// 	if j+1 < len(p) {
// 		// then check if the next letter is *
// 		if string(p[j+1]) == "*" {
// 			//   when we encounter a "*" we have 2 choice
// 			// one chose *
// 			// when we choose we move j +2 and when we dont we move i+1
// 			// fmt.Println("found *")
// 			// fmt.Println("creating 2 check on left(dont use * and jump j by2) and right(use * and increase i by 1)")
// 			leftCheck := rcheck(i, j+2, s, p, iterations)
// 			//  since right check is dependent on match we dont need to check at all if it doesnt match
// 			// we can save unnecesary run doing that.
// 			rightCheck := match && rcheck(i+1, j, s, p, iterations)
// 			// rightCheck = match && rcheck(i+1, j, s, p, iterations)

// 			// fmt.Println("right check: ", rightCheck)
// 			// fmt.Println("left check: ", leftCheck)
// 			return leftCheck || rightCheck
// 			// return (rcheck(i, j+2, s, p, iterations) || (match && rcheck(i+1, j, s, p, iterations)))
// 		}
// 	}
// 	//  here j is still in bound and  is still in bound, and also the next letter isnt "*"
// 	// then for this run, if theres a match we just check for the next letters in p and s i.e increment i,j by 1
// 	if match {
// 		fmt.Println("no * and match, moving up in both")
// 		return rcheck(i+1, j+1, s, p, iterations)
// 	}
// 	return false
// }

// func isMatch(s string, p string) bool {

// 	if p == ".*" {
// 		return true
// 	}
// 	// when p doesnt contain . and * its string matching
// 	if !strings.Contains(p, ".") && !strings.Contains(p, "*") {
// 		return s == p
// 	}
// 	sBytes := []byte(s)
// 	pBytes := []byte(p)

// 	// if we get here, we need to handle "*"
// 	// ab*
// 	// aab
// 	// steps, check if s[i] == p[i]
// 	// abbb
// 	// abc

// 	// ismatch := false
// 	i, j := 0, 0
// 	iterations := 0
// 	match := rcheck(i, j, sBytes, pBytes, &iterations)
// 	// fmt.Println("got result, ran in iteration(s) ", iterations)
// 	return match

// }

// func isMatch(s string, p string) bool {

// 	if p == ".*" {
// 		return true
// 	}
// 	// when p doesnt contain . and * its string matching
// 	if !strings.Contains(p, ".") && !strings.Contains(p, "*") {
// 		return s == p
// 	}
// 	sBytes := []byte(s)
// 	pBytes := []byte(p)
// 	// if we get here, we need to handle "*"
// 	// ab*
// 	// aab
// 	// steps, check if s[i] == p[i]
// 	// abbb
// 	// abc

// 	// ismatch := false
// 	i, j := 0, 0
// 	iterations := 0
// 	dp := map[string]bool{}
// 	match := rcheck(i, j, sBytes, pBytes, &iterations, dp)
// 	// fmt.Println("got result, ran in iteration(s) ", iterations)
// 	return match

// }
