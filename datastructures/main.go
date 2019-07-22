package main

import "fmt"

func main() {
	//m := []int{5,7,4,2,8,1,6}
	//printArray(m)
	m := []int{1, 2, 3, 4}
	k := 3
	Split(m, k)
	// segs := make([][]int, k)
	// res := make([][][]int, 0)
	// res = possibleSegments(m, k, segs, res)
	// printResult(res)
}

func printArray(data []int) {
	for i := 0; i < len(data); i++ {
		for j := 0; j <= i; j++ {
			print(data[j])
		}
		println()
	}
}

func possibleSegments(data []int, k int, segs [][]int, res [][][]int) [][][]int {
	if len(data) < k {
		return res
	}

	if k == 1 {
		segs[k-1] = data
		res = append(res, append(segs[:0:0], segs...))
		return res
	}

	for i := 1; i < len(data); i++ {
		segs[k-1] = data[:i]
		res = possibleSegments(data[i:], k-1, segs, res)
	}

	return res
}

func printResult(data [][][]int) {
	for _, v := range data {
		printSegs(v)

	}
}

func printSegs(data [][]int) {
	for _, v := range data {
		printArray1(v)

	}
	println("###################")
}

/*
	[1][2][3,4,5]
	[1][2,3][4,5]
	[1][2,3,4][5]
	[1,2][3][4,5]
	[1,2][3,4][5]
	[1,2,3][4][5]
*/

func printArray1(data []int) {
	str := "{"
	for _, v := range data {
		str += fmt.Sprintf("%d,", v)
	}
	print(str + "}")
}

func Split(data []int, k int) {
	for i := 0; i < len(data); i++ {
		s1 := append(data[:0:0], data[:i+1]...)
		// l := k
		for j := i + 1; j < len(data); j++ {
			s2 := append(data[:0:0], data[i+1:j+1]...)
			s3 := append(data[:0:0], data[j+1:]...)
			printArray(s1)
			printArray(s2)
			printArray(s3)

		}
	}
}
