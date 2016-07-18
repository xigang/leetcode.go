package longest_common_prefix

import (
	"testing"
)

func TestLongetCommonPrefix(t *testing.T) {
	strs := []string{"abab", "aba", "abc"}
	t.Log(LongestCommonPrefix(strs))
}
