package main

import "fmt"


func main() {
	array := []int{2, 4, 6, 8, 10}
	fmt.Println(binarySearch(array, 6))
}

func binarySearch(array []int, item int) int { // input: integer array, integer to search, return integer
	low := 0 // - low value
	high := len(array) - 1 // - high value - array length

	for low <= high { // while min value lesser or equals high

		mid := (low + high) / 2 // middle value - quotient of the sum of min and max values and divisor 2
		guess := array[mid] // guess value - value at the index of the middle value

		if guess == item { // the guess value is equal to the number we are looking for, then
			return mid // return middle value
		}

		if guess > item { // the guess value is greater than this number
			high = mid - 1 // the maximum value is middle - 1
		} else { // lesser
			low = mid + 1 // minimal equals middle + 1
		}
	}

	return -1 // else return -1 (item is not found)
}