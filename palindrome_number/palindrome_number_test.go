package palindrome_number

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	values := []struct {
		str int
		b   bool
	}{
		{1, true},
		{-1, false},
		{121, true},
		{16461, true},
	}

	for _, v := range values {
		if v.b != IsPalindrome(v.str) {
			t.Log(v.str, "failed")
		}
	}

}
