package main

import "fmt"

func reverseLeftWords(s string, n int) string {
	s = s[n:] + s[0:n]
	return s
}

func reverseLeftWords1(s string, n int) string {
	ns := ""
	for i := n; i < len(s); i++ {
		ns += s[i : i+1]
	}
	for i := 0; i < n; i++ {
		ns += s[i : i+1]
	}
	return ns
}
func main() {
	s := "12345"
	fmt.Println(reverseLeftWords(s, 2))
}
