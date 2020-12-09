package main

import "fmt"

/*
情况太多，考虑不过来，暂时放弃
*/

func isMatch(s string, p string) bool {
	// 如果字符串长度为0，需要检测下正则串
	if len(s) == 0 {
		// 如果正则串长度为奇数，必定不匹配，比如 "."、"ab*",必须是 a*b*这种形式，*在奇数位上
		if len(p)%2 != 0 {
			return false
		}
		i := 1
		for i < len(p) {
			if (p)[i] != '*' {
				return false
			}
			i += 2
		}
		return true
	}
	// 如果字符串长度不为0，但是正则串没了，return false
	if len(p) == 0 {
		return false
	}
	return dpMatch(&s, &p, 0, 0)
}

func dpMatch(s, p *string, i, j int) bool {

	if i >= len(*s) {
		if j >= len(*p) { //都匹配到头了
			return true
		} else if (*p)[j] == '*' { //p没有匹配到头，判断是否是*的情况
			if j == len(*p)-1 { //如果最后一位是*，则返回true
				return true
			} else { //最后还有其他字符，s到头了
				return false
			}
		}

	}

	//最好的情况
	if (*s)[i] == (*p)[j] { //a==a
		return dpMatch(s, p, i+1, j+1)
	} else {
		if (*p)[j] == '.' { //aa==.a
			return dpMatch(s, p, i+1, j+1)
		}
		if (*p)[j] == '*' {
			if (*s)[i] == (*p)[j-1] { //aaa=a*
				return dpMatch(s, p, i+1, j)
			} else { //aab=a*
				return dpMatch(s, p, i, j+1)
			}
		}
		//匹配任意字符的情况
		if (*p)[j+1] == '*' && j < len(*p) { //a==c*a
			return dpMatch(s, p, i, j+2)
		}

	}
	return false
}

func main() {
	s := "aa"
	p := "a*"
	fmt.Println(isMatch(s, p))
}
