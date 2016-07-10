package reverse_integer

import (
	"testing"
)

func TestReverse(t *testing.T) {
	values := []struct {
		n1 int
		n2 int
	}{
		{123, 321},
		{-123, -321},
		{900000, 9},
		{1463847412, 2147483641},
		{2147483647, 0},
		{2147483648, 0},
	}

	for _, v := range values {
		if v.n2 != Reverse(v.n1) {
			t.Log("failed.")
		}
	}
}
