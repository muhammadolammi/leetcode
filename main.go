package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func convert(s string, numRows int) string {
	if len(s) == 1 || s == "" || numRows == 1 {
		return s
	}
	returnString := ""

	stringLists := make([]string, numRows)
	direction := -1
	currentRow := 0
	for _, runeValue := range s {
		fmt.Println("running. row ", currentRow)
		stringLists[currentRow] += string(runeValue)
		if currentRow == 0 || currentRow == numRows-1 {
			direction *= -1
		}
		currentRow += direction
	}
	// for _, value := range stringLists {
	// 	returnString += value
	// }
	return returnString

}
func reverse(x int) int {
	isNegative := false
	xStringed := strconv.Itoa(x)
	if strings.Contains(xStringed, "-") {
		trimmedS := xStringed[1:]
		xStringed = string(trimmedS)
		isNegative = true
	}
	reversedStringed := ""
	for i := len(xStringed) - 1; i >= 0; i-- {
		reversedStringed += string(xStringed[i])
	}

	x, err := strconv.Atoi(reversedStringed)
	if err != nil {
		fmt.Println("converting error: ", err)
		return x
	}
	if x < math.MinInt32 {
		return 0
	}
	if x > math.MaxInt32 {
		return 0
	}

	if isNegative {
		x = -x
	}

	return x
}

func ap(a, b, k int) int {

	return a + ((k - 1) * b)
}

func main() {
	// fmt.Println(convert("AB", 1))
	// fmt.Println(reverse(-123))
	// fmt.Println(myAtoi("-2000000000000"))
	// fmt.Println(math.Pow10(20))
	//  ac*deee*f
	// adeee
	// fmt.Println(isPalindrome(-121))
	// n := 20
	// fmt.Println(sum(n))
	// fmt.Println((n * (n + 1) / 2))
	// fmt.Println(partition(4, 4))
	// dp := map[int]int{}
	// fmt.Println(fibonacci(9, dp))

	// fmt.Println(isMatch("acd", "ab*c*d"))
	// fmt.Println(an(1, 2, 5))

	// ls := []int{1, 5, 6, 20, 9, 14, 15, 3, 40, 33, 21, 32}
	// ls2 := []int{101, 200, 300, 400, 500, 500, 600, 200}
	// bubbleSort(&ls)
	// fmt.Println(ls)
	// QuickSort(&ls)
	// insertionSort(&ls2)

	// fmt.Println(ls)
	// fmt.Println(ls2)
	// s := MakeStack[int]()
	// s.Push(1)
	// s.Push(4)
	// s.Pop()
	// s.Push(5)
	// s.Push(10)
	// fmt.Println(s.Peek())

	// fmt.Println(s.Data)
}
