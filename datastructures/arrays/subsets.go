package arrays

/*
{1,2,3}

{1}, {2}, {3}
{1,2} {3}

*/

func Subsets(data []int) [][]int {
	subset := make([]int, 0)
	subsets := make([][]int, 0)
	return helper(data, subset, 0, subsets)
}

func helper(data []int, subset []int, index int, result [][]int) [][]int {
	if index >= len(data) {
		result = append(result, subset)
		return result
	}

	result = helper(data, subset, index+1, result)
	subset = append(subset, data[index])
	result = helper(data, subset, index+1, result)
	return result
}

// func printArray(data []int) {
// 	str := ""
// 	for _, v:= range data  {
// 		str += fmt.Sprintf("%d,", v)
// 	}
// 	fmt.Println(str)
// }
