package leetcode36

//尴尬，只是想用大神的record相似的方法写出来，结果一次过，再看大神，他竟然不用record这些空间换时间的方法...
//更尴尬的是，我发现空间换时间的方法，绝对比大神的快...100%...
func isValidSudoku(board [][]byte) bool {
	var row [9][10]bool
	var col [9][10]bool
	var block [9][10]bool
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '0'

				//只要有一个true，说明这行/列已经出现过一次
				// block[i/3*3+j/3][num]
				// i/3*3+j/3 代表第几个宫 建议看图理解，只可意会不可言传系列
				//
				// 九宫格
				// 1 2 3
				// 4 5 6
				// 7 8 9
				// 左到右是+1 所以j/3
				// 上到下是+3 所以i/3*3

				if row[i][num] || col[j][num] || block[i/3*3+j/3][num] {
					return false
				} else {
					row[i][num] = true
					col[j][num] = true
					block[i/3*3+j/3][num] = true
				}
			}
		}
	}
	return true
}
