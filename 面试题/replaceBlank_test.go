package 面试题

import (
	"fmt"
	"testing"
)

func TestReplaceBlank(t *testing.T){
	s := "sdajfkl dsfj sdfdo dfas dsf "
	if s1, ok := replaceBlank(s);ok{
		fmt.Println(s1)
	}else{
		fmt.Println("err")
	}
}
