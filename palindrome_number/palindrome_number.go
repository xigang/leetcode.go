/**
 * Determine whether an integer is a palindrome. Do this without extra space.
 */

package palindrome_number

import (
	"strconv"
)

func IsPalindrome(x int) bool {
	str := strconv.Itoa(x)
	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}
