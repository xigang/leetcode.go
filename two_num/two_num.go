/*Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution.

Example:
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].*/

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
