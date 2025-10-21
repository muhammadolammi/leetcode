package main

func fibonacci(n int, dp map[int]int) int {
	// dp := make(map[int]int,)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	cached, ok := dp[n]
	if ok {
		return cached
	}
	// if not cached, compute for n and store
	dp[n] = fibonacci(n-1, dp) + fibonacci(n-2, dp)
	return dp[n]

}

func sum(n int) int {
	if n <= 0 {
		return 0
	}
	// dp := make([]int, n+1)

	return n + sum(n-1)
}

func paths(n, m int) int {
	// write a function that take n , m and output the number of unique path from the top left corner
	// to bottom right corner of a NxM grid
	// Constaint you can only move down or right one step at a time

	// Solution
	// find the base case
	// for n =0 or m= 0 return 0 (no grid)
	// for n=1 you can only move left so one unique path. return 1
	// - -
	//for m=1 you can only move downward . so one unique path . return 1
	// -
	// -
	// conclusion for n==1 or m==1 return 1
	if n == 0 || m == 0 {
		return 0
	}
	if n == 1 || m == 1 {
		return 1
	}
	// Lets get recursive paterm
	// first lets play around with cases
	// paths (1,1)=1
	//paths (2, 1)=1
	// paths (1,2)=1
	// paths (2,2) = 2
	// - -
	// - -
	// paths (3,2) = 3
	// - -
	// - -
	// - -
	// paths (2, 3) = 3
	// - - -
	// - - -
	// lets do distance number
	// paths (3 ,5) = 15
	// - - - - -
	// - - - - -   5 4 3 2 1
	// - - - - -
	//  paths (2, 5)= 5
	// - - - - -
	// - - - - -
	//  paths(3, 4) = 10
	// - - - -
	// - - - -
	// - - - -
	// i noticed an intresting pattern , if either is 2, then the answer is the other
	// this may make it faster when we encounter 2
	if n == 2 {
		return m
	}
	if m == 2 {
		return n
	}

	return paths(n-1, m) + paths(n, m-1)

}

func partition(n, m int) int {
	//  write a func that counts the number of ways you can partition n objects usings parts up to m assume (m>=0)
	//  when given n objects see how many ways we can partition them up to m
	//  partition(0,0)=0 i.e arange empty object in zero partition , n=0 means include no part
	// partition(0, x) always equals 1. because we always include no part which on it own is a partition

	//  when m=0 and n is valid? we havee some objects/parts but no way to represent them
	//  partition(x, 0) = 0
	// if we get here n is already valid, so lets check m and retrun 0

	// lets test senarios
	// parts(1,1)
	// -
	// 1
	// parts (2, 1)
	// 1 + 1
	// parts(6, 1)      1
	// 1+1+1+1+1+1
	// parts (6, 2)     4
	// 2+2+2
	// 2+2+1+1
	// 2 + 1+ 1+1+1+1
	// 1+ 1+1+1+1+1

	// parts(6, 3)      7
	// 3+1+1+1
	// 3+3
	// 3+2+1
	// 2+2+2
	// 2+2+1+1
	// 2 + 1+ 1+1+1+1
	// 1+ 1+1+1+1+1
	// parts (6,4)        9
	// 4+2
	// 4+1+1
	// 3+1+1+1
	// 3+3
	// 3+2+1
	// 2+2+2
	// 2+2+1+1
	// 2 + 1+ 1+1+1+1
	// 1+ 1+1+1+1+1
	// parts (6,5)        10
	// 5+1
	// 4+2
	// 4+1+1
	// 3+1+1+1
	// 3+3
	// 3+2+1
	// 2+2+2
	// 2+2+1+1
	// 2 + 1+ 1+1+1+1
	// 1+ 1+1+1+1+1

	// parts (6,6)        11
	// 6
	// 5+1
	// 4+2
	// 4+1+1
	// 3+1+1+1
	// 3+3
	// 3+2+1
	// 2+2+2
	// 2+2+1+1
	// 2 + 1+ 1+1+1+1
	// 1+ 1+1+1+1+1
	// when m > n
	// parts (6,7)
	// 6
	// 5+1
	// 4+2
	// 4+1+1
	// 3+1+1+1
	// 3+3
	// 3+2+1
	// 2+2+2
	// 2+2+1+1
	// 2 + 1+ 1+1+1+1
	// 1+ 1+1+1+1+1
	// nothing for 7 so we can just do if m > n return partition(n,n)
	// managee negative input
	if n < 0 {
		return 0
	}
	if m > n {
		return partition(n, n)
	}
	if m == 1 || n == 1 || n == 0 {
		return 1
	}
	if m <= 0 {
		return 0
	}
	// the only thing that should really chnage per step is m
	return partition(n-m, m) + partition(n, m-1)

}
