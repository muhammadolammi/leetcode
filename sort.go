package main

func quickSortPartition(input []int, start, end int) int {
	// it return the new index of the pivot
	// Set pivot to the element at index high
	// Set i to the index before low
	// For each index (j) from low to high:
	// 	//If the element at index j is less than the pivot:
	// 		//Increment i by 1
	// 		//Swap the element at index i with the element at index j
	// 	Swap the element to the right of i with the pivot element
	// Return the new index of the pivot element (the item in the "middle" of the partition)
	pivot := input[end]
	i := start - 1
	j := start
	for ; j <= end; j++ {
		if input[j] < pivot {
			i++
			input[i], input[j] = input[j], input[i]
		}
		j++
	}
	input[i+1], input[end] = pivot, input[i+1]
	return i + 1
}
func quickSort(input []int, low, high int) {
	if low < high {
		mid := quickSortPartition(input, low, high-1)
		quickSort(input[:mid], 0, mid)
		quickSort(input[mid:], mid, high-1)
	}

}
func QuickSort(inputAdd *[]int) {
	input := *inputAdd
	quickSort(input, 0, len(input))
}
func insertionSort(inputAdd *[]int) {
	input := *inputAdd

	// i to range through the input
	for i := range input {
		j := i
		for j > 0 && input[j-1] > input[j] {
			//   swap
			input[j-1], input[j] = input[j], input[j-1]
			j--
		}

	}

}
func mergeSort(input []int) []int {
	// 	Input: A, an unsorted list of integers

	// If the length of A is less than 2, it's already sorted so return it
	// Split the input array into two halves down the middle
	// Call merge_sort() twice, once on each half
	// Return the result of calling merge(sorted_left_side, sorted_right_side) on the results of the merge_sort() calls
	lenght := len(input)
	if lenght < 2 {
		return input
	}

	leftSplit, rightSplit := input[0:lenght/2], input[lenght/2:]
	sorted_left := mergeSort(leftSplit)
	sorted_right := mergeSort(rightSplit)
	return merge(sorted_left, sorted_right)
}
func merge(inputLeft, inputRight []int) []int {
	// Inputs: A and B. Two sorted lists of integers

	// Create a new final list of integers.
	// Set i and j equal to zero. They will be used to keep track of indexes in the input lists (A and B).
	// Use a loop to compare the current elements of A and B. If an element in A is less than or equal to its respective element in B, add it to the final list and increment i. Otherwise, add the item in B to the final list and increment j. Continue until all items from one of the lists have been added.
	// After comparing all the items, there may be some items left over in either A or B. Add those extra items to the final list.
	// Return the final list.
	finalLs := []int{}
	leftLen := len(inputLeft)
	rightLen := len(inputRight)
	i, j := 0, 0
	for i < leftLen || j < rightLen {
		// if both are in bound
		if i < leftLen && j < rightLen {
			if inputLeft[i] < inputRight[j] {
				finalLs = append(finalLs, inputLeft[i])
				i++
				continue
			}
			if inputLeft[i] > inputRight[j] {
				finalLs = append(finalLs, inputRight[j])
				j++
				continue
			}
		}
		// all left are done but right contain some value
		if j < rightLen {
			remains := inputRight[j:]
			finalLs = append(finalLs, remains...)
			break
		}
		// all right are done but left contain some value

		if i < leftLen {
			remains := inputLeft[i:]
			finalLs = append(finalLs, remains...)
			break
		}
	}

	return finalLs

}
func bubbleSort(inputAdd *[]int) {
	//  keep swapping all elements from left to right untill theres nothing to swap
	// 	Set swapping to True
	// Set end to the length of the input list
	// While swapping is True:
	// Set swapping to False
	// For i from the 2nd element to end:
	// If the (i-1)th element of the input list is greater than the ith element:
	// Swap the (i-1)th element and the ith element
	// Set swapping to True
	// Decrement end by one
	// Return the sorted list
	input := *inputAdd
	swap := true
	end := len(input)
	for swap {
		swap = false
		for i := range input[:end] {
			if i == 0 {
				continue
			}
			if input[i-1] > input[i] {
				input[i-1], input[i] = input[i], input[i-1]
				swap = true
			}

		}
		end--
	}
}
