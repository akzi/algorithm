package kmp

import (
	"math/rand"
	"reflect"
	"testing"
)

func NMSearch(str, subStr string) []int {
	var (
		SLen   = len(str)
		subLen = len(subStr)
		result []int
		i, j   = 0, 0
	)
	if SLen == 0 || subLen == 0 {
		return result
	}
	for i < SLen {
		for i < SLen && j < subLen {
			if str[i] == subStr[j] {
				i++
				j++
			} else {
				i = i - j + 1
				break
			}
		}
		if j == subLen {
			result = append(result, i-j)
		}
		j = 0
	}
	return result
}

func generateStr() string {
	var str string
	for len(str) < 32 {
		str += string([]byte{byte(rand.Intn(3) + 'a')})
	}
	return str
}

func TestSearch(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		str := generateStr()
		sub := str[1:3]
		if reflect.DeepEqual(NMSearch(str, sub), Search(str, sub)) == false {
			t.Error(str, sub, NMSearch(str, sub), Search(str, sub))
		}
	}

}
func TestGetNext(t *testing.T) {
	subStr := "aabb"
	str := "aabaabaaaabbaabaabaaabbaabaabb"
	indexs := Search(str, subStr)
	if len(indexs) != 3 {
		t.Error(indexs)
	}
	for _, i := range indexs {
		if str[i:i+len(subStr)] != subStr {
			t.Error(i)
		}
	}
}
