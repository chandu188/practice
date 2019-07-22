package dynamicprogramming

/*
Given a “2 x n” board and tiles of size “2 x 1”,
count the number of ways to tile the given board using the 2 x 1 tiles.
A tile can either be placed horizontally i.e., as a 1 x 2 tile or vertically i.e., as 2 x 1 tile.



*/

func TilingProblem(n int) int {
	return count2(n, 1, 1)
}

func count(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	return count(n-2) + count(n-1) // same as fibnocci
}

func count2(n int, a, b int) int {
	if n == 0 {
		return a
	}
	if n == 1 {
		return b

	}

	return count2(n-1, b, a+b)
}
