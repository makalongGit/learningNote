package 面试题

import "strings"

/**
请实现一个算法，确定一个字符串的所有字符【是否全都不同】。这里我们要求【不允许使用额外的存储结构】。
给定一个string，请返回一个bool值,true代表所有字符全都不同，false代表存在相同的字符。
保证字符串中的字符为【ASCII字符】。字符串的长度小于等于【3000】。
*/

//通过String.Count 函数

func isUniqueString1(s string) bool {
	if len(s) > 3000 {
		return false
	}
	for _, v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}

//通过index 和lastIndex判断
func isUniqueString2(s string) bool {
	if len(s) > 3000 {
		return false
	}
	for _, v := range s {
		if v > 127 {
			return false
		}
		if strings.Index(s, string(v)) != strings.LastIndex(s, string(v)) {
			return false
		}
	}
	return true
}
