package stack

import "data_struct/ArrayStack"

func IsValid(s string) bool {
	size := len(s)
	stack := ArrayStack.Constructor(size)
	for i := 0; i < size; i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack.Push(s[i])
		} else {
			if stack.GetSize() == 0 {
				return false
			}
			top := stack.Pop()
			var m byte
			if s[i] == byte(')') {
				m = byte('(')
			} else if s[i] == byte('}') {
				m = byte('{')
			} else {
				m = byte('[')
			}
			if m != top {
				return false
			}
		}
	}
	if stack.GetSize() != 0 {
		return false
	}
	return true
}

//利用字节编码
func IsValid1(s string) bool {
	size := len(s)
	top := 0
	stack := make([]byte, size)
	for i := 0; i < size; i++ {
		c := s[i]
		switch c {
		case '(':
			stack[top] = c + 1
			top++
		case '[', '{':
			stack[top] = c + 2
			top++
		case ']', '}', ')':
			if top > 0 && stack[top-1] == c {
				top--
			} else {
				return false
			}
		}
	}
	return top == 0
}
