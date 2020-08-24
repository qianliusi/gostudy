package leetcode

import (
	"bytes"
)

func replaceSpace(s string) string {
	var buffer bytes.Buffer
	for _, v := range s {
		if v == ' ' {
			buffer.WriteString("%20")
		} else {
			buffer.WriteRune(v)
		}
	}
	return buffer.String()
}
