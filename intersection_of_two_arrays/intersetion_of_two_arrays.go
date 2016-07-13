/**

Given two arrays, write a function to compute their intersection.

Example:
Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2].

Note:
Each element in the result must be unique.
The result can be in any order.

*/

package intersection_of_two_arrays

func Intersection(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]int)
	for index, value := range nums1 {
		nums1Map[value] = index
	}

	nums2Map := make(map[int]int)
	for index, value := range nums2 {
		nums2Map[value] = index
	}

	var ret []int
	for x, _ := range nums2Map {
		for y, _ := range nums1Map {
			if x == y {
				ret = append(ret, x)
			}
		}
	}
	return ret
}
