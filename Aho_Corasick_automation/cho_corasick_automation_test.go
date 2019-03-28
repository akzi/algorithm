package aho_corasick_automation

import (
	"fmt"
	"testing"
)

func TestAhoCorasickAutomation_Search(t *testing.T) {
	corasickAutomation := NewAhoCorasickAutomation()
	corasickAutomation.Append("abcd")
	corasickAutomation.Append("bcd")
	corasickAutomation.Append("cd")
	corasickAutomation.Append("d")

	corasickAutomation.Append("1")
	corasickAutomation.Append("12")
	corasickAutomation.Append("123")
	corasickAutomation.Append("1234")

	corasickAutomation.Build()

	target := "012345"
	result := corasickAutomation.Search(target)
	if len(result) != 4 {
		t.Fatal(result)
	}
	for _, it := range result {
		fmt.Println(target[it[0] : it[0]+it[1]])
	}
}
