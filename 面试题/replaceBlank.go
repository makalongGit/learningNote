package 面试题

import (
	"strings"
	"unicode"
)

/**
字符串替换问题
问题描述

请编写一个方法，将字符串中的空格全部替换为“%20”。
假定该字符串有足够的空间存放新增的字符，并且知道字符串的真实长度(小于等于1000)，同时保证字符串由【大小写的英文字母组成】。
给定一个string为原始的串，返回替换后的string。
 */

func replaceBlank(s string) (string,bool){
	rune_s := []rune(s)
	if len(rune_s) > 1000 {
		return s,false
	}
	for _,v := range s{
		if unicode.IsLetter(v) == false && string(v) != " " {
			return s,false
		}
	}
	return strings.Replace(s, " ", "%20", -1), true
}