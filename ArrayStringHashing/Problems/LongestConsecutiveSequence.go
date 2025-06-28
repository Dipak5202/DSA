package problems

func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longest := 0

	for num := range numSet {
		// Only start if num-1 is not in the set
		if _, exists := numSet[num-1]; !exists {
			currentNum := num
			streak := 1

			for {
				_, ok := numSet[currentNum+1]
				if !ok {
					break
				}
				currentNum++
				streak++
			}

			if streak > longest {
				longest = streak
			}
		}
	}

	return longest
}
