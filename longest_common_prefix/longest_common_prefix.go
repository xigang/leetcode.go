/**

 Write a function to find the longest common prefix string amongst an array of strings.

Subscribe to see which companies asked this question
*/

package longest_common_prefix

func LongestCommonPrefix(strs []string) string {
	var word string
	if len(strs) <= 0 {
		return word
	}

	for i := 1; i <= len(strs[0]); i++ {
		w := strs[0][:i]
		match := true
		for j := 1; j < len(strs); j++ {
			if i > len(strs[j]) || w != strs[j][:i] {
				match = false
				break
			}
		}

		if !match {
			return word
		}

		word = w
	}

	return word
}
