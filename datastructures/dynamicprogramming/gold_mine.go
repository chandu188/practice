package dynamicprogramming


func GoldMine(mine [][]int, m, n int) int{
	 dp := make([][]int, len(mine))
	 for i, r := range mine {
		 dp[i] = make([]int, len(r))
	 }

	 for col:=n-1; col >=0; col-- {
		 for row:=0; row<m; row ++ {
			var right, rightUp, rightDown int
			if col == n-1 {
				right = 0
			} else {
				right = dp[row][col+1]
			}

			if col==n-1 || row ==0 {
				rightUp = 0 
			}else{
				rightUp = dp[row-1][col+1]
			}
			
			if col==n-1 || row==m-1 {
				rightDown = 0
			} else{
				rightDown = dp[row+1][col+1]
			}
			dp[row][col] = mine[row][col] + max(right, max(rightUp, rightDown)) 
		 }
	 }
	 
	 max := -1
	 for _, val := range dp{
		 if max < val[0] {
			 max = val[0]
		 }
	 }

	 return max

}

func max(a, b int) int {
	if a > b {
		return a 
	}
	return b
}

func _terneray(cond bool, res1, res2 int) int {
	if cond {
		return res1
	}

	return res2
}