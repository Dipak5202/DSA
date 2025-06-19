package selectionsort

func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		mindex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[mindex] > arr[j] {
				mindex = j
			}
		}
		if mindex != i {
			temp := arr[i]
			arr[i] = arr[mindex]
			arr[mindex] = temp
		}
	}
}
