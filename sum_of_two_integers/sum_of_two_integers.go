/**
Calculate the sum of two integers a and b, but you are not allowed to use the operator + and -.
Example:
Given a = 1 and b = 2, return 3.
*/

package sum_of_two_integers

func GetSum(a int, b int) int {
	var c int
	for b != 0 {
		c = a & b
		a ^= b
		b = c << 1
	}
	return a
}
