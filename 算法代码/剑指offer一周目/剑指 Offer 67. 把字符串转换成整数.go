package main

import (
	"fmt"
	"strings"
)

//100,20
func strToInt(str string) int {
	str = strings.Trim(str, " ")
	l := len(str)
	if l == 0 {
		return 0
	}
	isNeg := false
	if str[0] == '-' {
		isNeg = true
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}
	//str=strings.TrimLeft(str[i:],"0")
	for len(str) > 0 && str[0] == '0' {
		str = str[1:]
	}
	l = len(str)
	if l == 0 {
		return 0
	}
	//最大只有10位,所以取11位
	if l > 10 {
		str = str[:11]
		l = 11
	}
	//使用int64主要是考虑10位的情况会有大于2^13-1的
	var res int64

	for i := 0; i < l; i++ {
		if str[i] < '0' || str[i] > '9' {
			break
		} else {
			res *= 10
			res += int64(str[i] - '0')
		}
	}
	if isNeg && res > (1<<31) {
		return -1 << 31
	} else if !isNeg && res > 1<<31-1 {
		return 1<<31 - 1
	} else {
		if isNeg {
			return -int(res)
		}
		return int(res)
	}
	//s,_:=strconv.Atoi(str)
	//return s
}
func main() {
	s := "4193 with words"
	fmt.Println(strToInt(s))
}
