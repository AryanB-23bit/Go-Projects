//Insertion Sort implementation in Go

package main

import "fmt"

func main() {
	//Test
	fmt.Println(insertionSort([]int{5,1,3,7,6}))

	//Test
	fmt.Println(insertionSort([]int{41,51,69,36,51,68}))
}

func insertionSort(unsortedList []int) []int {
	for index := 1; index < len(unsortedList); index++ {
		for key:= index; key > 0 && unsortedList[key -1] > unsortedList[key]; {
			unsortedList[key - 1], unsortedList[key] = unsortedList[key], unsortedList[key - 1]
		key = key - 1
		}
	}
	sortedList := unsortedList
	return sortedList
}
