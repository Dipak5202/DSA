package problems

import "sort"

type kv struct {
	Key   int
	Value int
}

func TopKFrequent(nums []int, k int) []int {
	result := make([]int, k)
	check := make(map[int]int)
	for _, ele := range nums {
		check[ele]++
	}

	var freqPairs []kv
	for key, value := range check {
		freqPairs = append(freqPairs, kv{key, value})
	}

	sort.Slice(freqPairs, func(i int, j int) bool {
		return freqPairs[i].Value > freqPairs[j].Value
	})

	for i := 0; i < k && i < len(freqPairs); i++ {
		result[i] = freqPairs[i].Key
	}

	return result
}
