package main

import (
	"fmt"
	"strings"
)

//库函数
//100,51
func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	ss := strings.Split(s, " ")
	sb := new(strings.Builder)
	strings.TrimSpace(s)
	for i := len(ss) - 1; i >= 0; i-- {
		fmt.Println(ss[i])
		if ss[i] != "" {
			sb.WriteString(ss[i])
			if i != 0 {
				sb.WriteString(" ")
			}
		}
	}
	return sb.String()
}

//双指针
//100,50
func reverseWords1(s string) string {
	l := len(s) - 1
	if l == -1 {
		return ""
	}
	//去掉首尾空格，方便后续处理
	s = strings.Trim(s, " ")
	l = len(s) - 1
	r := l
	sb := new(strings.Builder)
	for l >= 0 {
		//从尾部开始找单词的第一位
		for l >= 0 && s[l] != ' ' {
			l--
		}
		//写入
		sb.WriteString(s[l+1:r+1] + " ")

		//跳过空格，找下一个单词的结尾
		for l >= 0 && s[l] == ' ' {
			l--
		}
		//更新r，现在l和r都在第二个单词的结尾处等待下一轮循环
		r = l
	}

	//去掉多余的最后一个空格
	return strings.Trim(sb.String(), " ")
}

func main() {
	s := "  the  world is   beautiful   "
	fmt.Println(reverseWords1(s))
}
