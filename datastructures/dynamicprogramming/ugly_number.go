package dynamicprogramming

// 1,2,3,4,5,6,8,10...
func UglyNumber(n int) int{
	i2:= 0
	i3:= 0
	i5:= 0
	next_multiple_2 := 2
	next_multiple_3 := 3
	next_multiple_5 := 5
	uglyNums := make([]int, 0)
	uglyNums =append(uglyNums, 1) 
	for i:=1;i<n;i++ {
		next_ugly_num := min(min(next_multiple_2,next_multiple_3), next_multiple_5)
		uglyNums =append(uglyNums, next_ugly_num)
		if next_ugly_num == next_multiple_2 {
			i2 = i2 +1
			next_multiple_2 = 2 * uglyNums[i2]
		} else if next_ugly_num == next_multiple_3{
			i3 = i3 +1
			next_multiple_3 = 3 * uglyNums[i3]
		} else if next_ugly_num == next_multiple_5{
			i5 = i5 + 1
			next_multiple_5 = 5 * uglyNums[i5]
		}
		
	}
	return uglyNums[n-1]
}

func min(a, b int)int{
	if a <= b {
		return a 
	}
	return b
}