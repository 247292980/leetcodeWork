package leetcode14

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}
	short := strs[0]
	for _, s := range strs {
		if len(short) > len(s) {
			short = s
		}
	}

	for i, r := range short {
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != byte(r) {
				return strs[j][:i]
			}
		}
	}
	return short
}
