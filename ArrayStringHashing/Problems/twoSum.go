package problems

// TwoSum finds two numbers in the array that add up to the target.
// It returns the indices of the two numbers.
func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	ans := make([]int, 2)
	for i, num := range nums {
		complement := target - num
		if j, found := numMap[complement]; found {
			ans[0] = j
			ans[1] = i
			return ans
		}
		numMap[num] = i
	}
	return ans
}

func twoSum(nums []int, target int) []int {
	anser := make([]int, 2)
	for i := 0; i < len(nums)-1; i++ {
		for j := i; j < len(nums)-1; j++ {
			if (nums[i] + nums[j+1]) == target {
				anser[0] = i
				anser[1] = j + 1
				return anser
			}
		}
	}
	return anser
}
