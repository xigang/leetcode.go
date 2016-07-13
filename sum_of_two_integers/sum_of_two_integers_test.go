package sum_of_two_integers

import (
	"testing"
)

func TestSGetSum(t *testing.T) {
	values := []struct {
		a, b, sum int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{4, 5, 9},
	}

	for _, v := range values {
		if v.sum == GetSum(v.a, v.b) {
			t.Log("{PASS}")
		}
	}
}
