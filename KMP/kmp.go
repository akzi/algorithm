package kmp

func GetNext(str string) []int {
	var (
		strLen = len(str)
		res    = make([]int, strLen)
		i, j   = 1, 0
	)
	for i < strLen-1 {
		if str[i] == str[j] {
			i++
			j++
			res[i] = j
		} else {
			j = res[j]
			if j == 0 {
				i++
			}
		}
	}
	return res
}

func Search(str, sub string) []int {
	var (
		i, j   = 0, 0
		result []int
		strLen = len(str)
		subLen = len(sub)
		next   = GetNext(sub)
	)

	for i < strLen {
		if str[i] == sub[j] {
			i++
			j++
		} else {
			if next[j] == 0 && j == 0 {
				i++
			} else {
				j = next[j]
			}
		}
		if j == subLen {
			result = append(result, i-j)
			j = 0
		}
	}
	return result
}
