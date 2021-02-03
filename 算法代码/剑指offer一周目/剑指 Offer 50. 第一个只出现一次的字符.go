package main

import "fmt"

//18,14
func firstUniqChar(s string) byte {
	l := len(s)
	if l == 0 {
		return ' '
	}
	dp := make(map[byte][2]int, 26)
	for i := 0; i < l; i++ {
		if v, ok := dp[s[i]]; ok {
			v[1] += 1
			dp[s[i]] = v
		} else {
			dp[s[i]] = [2]int{i, 1}
		}
	}
	for i := 0; i < l; i++ {
		if dp[s[i]][1] == 1 {
			return s[i]
		}
	}
	return ' '
}

//98,99
func firstUniqChar1(s string) byte {
	var list [26]int
	length := len(s)
	for i := 0; i < length; i++ {
		list[s[i]-'a']++
	}
	for i := 0; i < length; i++ {
		if list[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}

func main() {
	s := "llcczzxxzzaaxxxxsstt"
	fmt.Println(firstUniqChar(s))
}
