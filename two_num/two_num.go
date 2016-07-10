package twosum

func TwoSum(array []int, target int) []int {
	var result []int
	for i, _ := range array {
		for j := i + 1; j < len(array); j++ {
			if array[i]+array[j] == target {
				result = append(result, i, j)
			}
		}
	}
	return result
}
