/*
Implement atoi to convert a string to an integer.

Hint: Carefully consider all possible input cases. If you want a challenge, please do not see below and ask yourself what are the possible input cases.

Notes: It is intended for this problem to be specified vaguely (ie, no given input specs). You are responsible to gather all the input requirements up front.
*/

package string_to_integer

import (
	"math"
	"strings"
)

func MyAtoi(str string) int {
	if str == "" {
		return 0
	}

	var (
		neg    bool
		ret    int
		t1, t2 bool
	)
	for _, v := range strings.TrimSpace(str) {
		if string(v) == "-" {
			t1 = true
		}

		if string(v) == "+" {
			t2 = true
		}
	}

	if t1 && t2 {
		return 0
	}

	for _, v := range strings.TrimSpace(str) {
		if string(v) == "-" {
			neg = true
			continue
		} else if string(v) == "+" {
			continue
		}

		digit := int(v - '0')
		if neg {
			if -ret < (math.MinInt32+digit)/10 {
				return math.MinInt32
			}
		} else {
			if ret > (math.MaxInt32-digit)/10 {
				return math.MaxInt32
			}
		}
		ret = ret*10 + digit
	}

	if neg {
		ret = -ret
	}
	return ret
}
