package leetcode6

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var rune []rune = []rune(s)
	var zStr string
	var separatorLen = numRows*2 - 2
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < len(s); j += separatorLen {
			zStr = zStr + string(rune[j+i])
			if i != 0 && i != numRows-1 && j+separatorLen-i < len(s) {
				zStr = zStr + string(rune[j+separatorLen-i])
			}
		}
	}
	return zStr
}
