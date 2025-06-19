package main

import (
	bubblesorting "DSA/Sorting/BubbleSorting"
	mergesort "DSA/Sorting/MergeSort"
	selectionsort "DSA/Sorting/SelectionSort"
)

func main() {
	BubblesortingExample()
	SelectionSortExample()
	MergeSortExample()
}

func SelectionSortExample() {
	println()

	arr := []int{64, 34, 25, 12, 22, 11, 90}
	for i := 0; i < len(arr); i++ {
		print(arr[i], " ")
	}
	selectionsort.SelectionSort(arr)
	println()
	for i := 0; i < len(arr); i++ {
		print(arr[i], " ")
	}
	println()
}

func BubblesortingExample() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	for i := 0; i < len(arr); i++ {
		print(arr[i], " ")
	}
	bubblesorting.Bubblesorting(arr)
	println()
	for i := 0; i < len(arr); i++ {
		print(arr[i], " ")
	}
}

func MergeSortExample() {
	println()
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	for i := 0; i < len(arr); i++ {
		print(arr[i], " ")
	}
	sorted := mergesort.MergeSort(arr)
	println()
	for i := 0; i < len(sorted); i++ {
		print(sorted[i], " ")
	}
}
