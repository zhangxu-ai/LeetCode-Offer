package main

import (
	"fmt"
	"strings"
)

func replaceSpace(s string) string {
	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] == ' ' {
			s = s[:i] + "%20" + s[i+1:]
			l += 2
		}
	}
	return s
}

func replaceSpace1(s string) string {
	ss := strings.Split(s, " ")
	s = ""
	for i, s2 := range ss {
		s += s2
		if i == len(ss)-1 {
			continue
		} else {
			s += "%20"
		}
	}
	return s
}

func replaceSpace2(s string) string {
	var res strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			res.WriteString("%20")
		} else {
			res.WriteByte(s[i])
		}
	}
	return res.String()
}

func replaceSpace3(s string) string {
	s1 := strings.ReplaceAll(s, " ", "%20")
	return s1
}

func main() {
	s := "We are happy."
	s = replaceSpace3(s)
	fmt.Println(s)
}
