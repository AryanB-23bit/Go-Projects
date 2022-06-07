package main

import (
	"fmt"
)

func main() {
	listOne := []int {9,1,5,3,7}
	fmt.Println(mergeSort(listOne))

	listTwo := []int {3, 41, 52, 26, 38, 57, 9}
	fmt.Println(mergeSort(listTwo))
}


func mergeSort(numArr []int) []int{
	if len(numArr) <= 1{
		return numArr
	}

	mid := len(numArr) / 2

	leftHalf := mergeSort(numArr[:mid]) // in go :mid is not inclusive so it will only go till Arr[0] to Arr[mid - 1]
	rightHalf := mergeSort(numArr[mid:])

	return merge(leftHalf,rightHalf)

}

func merge(subArrayLeft []int, subArrayRight []int) []int {
	var sortedArr []int
	i, j := 0, 0
	for i < len(subArrayLeft) && j < len(subArrayRight) {
		if subArrayLeft[i] < subArrayRight[j] {
			sortedArr = append(sortedArr, subArrayLeft[i])
			i++
		} else {
			sortedArr = append(sortedArr, subArrayRight[j])
			j++
		}
	}
	for i < len(subArrayLeft){
		sortedArr = append(sortedArr, subArrayLeft[i])
		i++
	}
	for j < len(subArrayRight) {
		sortedArr = append(sortedArr, subArrayRight[j])
		j++
	}

	return sortedArr
}
