/*
Reverse digits of an integer.

Example1: x = 123, return 321
Example2: x = -123, return -321
*/

package reverse_integer

import (
	"math"
)

func Reverse(x int) int {
	var y, n int
	for x != 0 {
		n = x % 10
		if y > math.MaxInt32/10 || y < math.MinInt32/10 {
			return 0
		}
		y = y*10 + n
		x = x / 10
	}
	return y
}
