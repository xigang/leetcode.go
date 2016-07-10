package twosum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	var (
		// array1  = []int{3, 2, 4}
		// target1 = 6
		array1  = []int{0, 4, 3, 0}
		target1 = 0
	)
	result := TwoSum(array1, target1)
	t.Log(result)
}
