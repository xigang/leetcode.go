/**
Given two arrays, write a function to compute their intersection.

Example:
Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2, 2].

Note:
	Each element in the result should appear as many times as it shows in both arrays.
	The result can be in any order.
Follow up:
	What if the given array is already sorted? How would you optimize your algorithm?
	What if nums1's size is small compared to nums2's size? Which algorithm is better?
	What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?
*/

package intersection_of_two_arrays2

import (
	"sort"
)

func Intersect(nums1 []int, nums2 []int) []int {
	var ret []int

	sort.Ints(nums1)
	sort.Ints(nums2)
	var i, j int
	for i < len(nums1) && j < len(nums2) {
		n1 := nums1[i]
		n2 := nums2[j]
		if n1 == n2 {
			ret = append(ret, n1)
			i++
			j++
		} else if n1 < n2 {
			i++
		} else {
			j++
		}
	}
	return ret
}
