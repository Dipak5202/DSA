package main

import searching "DSA/Searching/BinarySerch"

func main() {

	arr := []int{23, 45, 67, 89, 234, 456, 678, 890, 1234, 5678}
	target := 890
	result := searching.BinarySearch(arr, target)
	if result == -1 {
		println("Element not found in the array")
	} else {
		println("Element found at index:", result)
	}
}
