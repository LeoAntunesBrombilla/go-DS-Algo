package main

import "fmt"

// O(n^2) time | O(1) space
func insertionSort(array []int) []int {
	for index := 1; index < len(array); index++ {
		for count := index; count > 0 && array[count-1] > array[count]; count-- {
			array[count-1], array[count] = array[count], array[count-1]
		}
	}
	return array
}

func main() {

	exemplo := []int{8, 5, 2, 9, 5, 6, 3}
	fmt.Println(insertionSort(exemplo))
}
