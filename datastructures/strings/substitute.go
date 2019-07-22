package strings

type Operation struct {
	startIndex int
	endIndex   int
	substitute string
}

func substitute(str string, ops []Operation) string {
	i := 0
	j := 0
	res := ""
	for i < len(str) {
		if j < len(ops) && i == ops[j].startIndex {
			res = res + ops[j].substitute
			i += (ops[j].endIndex - ops[j].startIndex)
			j++
		} else {
			res += string(str[i])
		}
		i++
	}
	return res
}
