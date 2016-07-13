package intersection_of_two_arrays

import (
	"testing"
)

//Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2].
func TestIntersection(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	ret := Intersection(nums1, nums2)
	t.Log(ret)
}
