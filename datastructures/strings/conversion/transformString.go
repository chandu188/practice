package conversion

/*
	Given two strings s1 and s2(call letters in uppercase). Check if it is possible to convert s1 to s2 by performing following operations.

1. Make some lowercase letters uppercase.
2. Delete all the lowercase letters.

Input : s1 = daBcd s2 = ABC
Output : yes
Explanation : daBcd -> dABCd -> ABC
Covert a and b at index 1 and 3 to
upper case, delete the rest those are
lowercase. We get the string s2.

*/

import "strings"

func TransfromString(a, b string) bool {
	strs := segments(a)
	for _, v := range strs {
		if strings.ToLower(v) == strings.ToLower(b) {
			return true
		}
	}
	return false
}

func segments(a string) []string {
	return helper(a, "", 0, make([]string, 0))
}

func helper(a string, sub string, ind int, res []string) []string {
	if ind >= len(a) {
		res = append(res, sub)
		return res
	}

	res = helper(a, sub, ind+1, res)
	sub = sub + string(a[ind])
	res = helper(a, sub, ind+1, res)
	return res
}
