package problems

import "sort"

func GroupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)
	for _, str := range strs {
		sorted := sortString(str)
		anagramMap[sorted] = append(anagramMap[sorted], str)
	}
	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}
	return result
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
