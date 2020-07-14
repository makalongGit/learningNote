package 面试题

import "strings"

/**
判断两个给定的字符串排序后是否一致
问题描述
给定两个字符串，请编写程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
这里规定【大小写为不同字符】，且考虑字符串重点空格。
给定一个string s1和一个string s2，请返回一个bool，代表两串是否重新排列后可相同。
保证两串的长度都小于等于5000。
*/
func isRegroup(s1 string, s2 string) bool {
	sln1 := len([]rune(s1))
	sln2 := len([]rune(s2))
	if sln1 > 5000 || sln2 > 5000 || sln1 != sln2 {
		return false
	}
	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)){
			return false
		}
	}
	return true
}
