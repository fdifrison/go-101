package iteration

import "strings"

func Repeat(s string, n int) string {
	var toRepeat string
	for range n {
		toRepeat += s
	}
	return toRepeat
}

func RepeatFast(s string, n int) string {
	var toRepeat strings.Builder
	for range n {
		toRepeat.WriteString(s)
	}
	return toRepeat.String()
}
