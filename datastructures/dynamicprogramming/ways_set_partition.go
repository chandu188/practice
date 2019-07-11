package dynamicprogramming

/*
	Input:  n = 2
Output: Number of ways = 2
Explanation: Let the set be {1, 2}
            { {1}, {2} }
            { {1, 2} }

Input:  n = 3
Output: Number of ways = 5
Explanation: Let the set be {1, 2, 3}
             { {1}, {2}, {3} }
             { {1}, {2, 3} }
             { {2}, {1, 3} }
             { {3}, {1, 2} }
             { {1, 2, 3} }.

	S(n,k) be the number of ways to partitions n numbers into k sets
	For n=2 k=1,2
		  s(2,1) = {{1,2}}
		  s(2,2) = {{1},{2}}
	For n=3 k=1,2,3
		  s(3,1) = {{1,2,3}}
		  s(3,2) = [{{1,2},{3}}
					{{1,3},{2}}
					{{2,3},{1}}]
		   s(3,3) = {{1},{2},{3}}

	s(3,2) = 2*s(2,2) (add {3} per each partiotion) + s(2,1) (add{3} to {{1,2}})
	s(n+1, k) = k*s(n,k) + s(n,k-1)

*/
func SetPartition(n, k int) int {
	if k == 1 {
		return 1
	}
	s := make([][]int, n+1)
	for i := range s {
		s[i] = make([]int, k+1)
	}

	for i := 1; i <= n; i++ {
		s[i][1] = 1
	}

	for i := 2; i <= n; i++ {
		for j := 2; j <= k; j++ {
			s[i][j] = j*s[i-1][j] + s[i-1][j-1]
		}
	}

	res := 0
	for i:=1;i<=k;i++ {
		res = res + s[n][i]
	} 
	return res
}
