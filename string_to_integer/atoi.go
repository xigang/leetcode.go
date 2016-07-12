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

	cs := strings.TrimSpace(str)
	var i int
	for ; i < len(cs); i++ {
		if string(cs[i]) == "+" || string(cs[i]) == "-" {
			if i == 0 {
				continue
			}
		}

		if int(cs[i]-'0') > 9 || int(cs[i]-'0') < 0 {
			return 0
		}
	}

	var neg, ret, index int
	neg = 1
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
