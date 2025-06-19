package main

import (
	"sort"
)

func main() {

	println("unsorted array:")

	arr := []int{64, 34, 25, 12, 22, 11, 90}
	for i := 0; i < len(arr); i++ {
		print(arr[i], " ")
	}
	println()

	//sort.Slice
	println("Using sort.Slice:")
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	println()
	for i := 0; i < len(arr); i++ {
		print(arr[i], " ")
	}
	println()
	println("End  Sort Example")

	//sort.SearchInts
	res := sort.SearchInts(arr, 22)
	println("Result of sort.SearchInts:", res)

	//sort.Find
	i, found := sort.Find(len(arr), func(i int) int {
		if arr[i] == 22 {
			return 0
		} else if arr[i] < 22 {
			return -1
		}
		return 1
	})
	println("Result of sort.Find:", i, found)

	//sort.SliceIsSorted
	println("Is the array sorted?", sort.SliceIsSorted(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	}))
	println("End of Sort Library Example")

	//sort.IsSorted
	println("Is the array sorted using sort.IsSorted?", sort.IsSorted(sort.IntSlice(arr)))

	//sort.IntsAreSorted
	println("Is the array sorted using sort.IntsAreSorted?", sort.IntsAreSorted(arr))

	println("End of Sort Library Example")

}
