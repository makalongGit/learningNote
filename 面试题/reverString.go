package 面试题

/**
请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。
给定一个string，请返回一个string，为翻转后的字符串。保证字符串的长度小于等于5000。
*/

func reverString(s string) (string, bool) {
	str := []rune(s)

	len := len(str)
	if len > 5000 {
		return "", false
	}

	for i := 0; i < len/2; i++ {
		str[i], str[len-i-1] = str[len-i-1], str[i]
	}
	return string(str), true
}
