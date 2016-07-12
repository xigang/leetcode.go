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

	var neg, ret, index int
	neg = 1
	cs := strings.TrimSpace(str)
	for _, v := range cs {
		if string(v) == "+" {
			index++
		} else if string(v) == "-" {
			neg = -1
			index++
		}
	}

	for ; index < len(cs); index++ {
		if int(cs[index]-'0') < 0 || int(cs[index]-'0') > 9 {
			break
		}
		ret = ret*10 + int(cs[index]-'0')
		if ret > math.MaxInt32 {
			break
		}
	}

	if neg*ret >= math.MaxInt32 {
		return math.MaxInt32
	} else if neg*ret <= math.MinInt32 {
		return math.MinInt32
	}
	return ret * neg
}
