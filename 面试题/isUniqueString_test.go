package 面试题

import (
	"fmt"
	"testing"
)

func TestIsUniqueString1_test(t *testing.T){
	words := "dfjkalsfu983"
	fmt.Println(isUniqueString1(words))
	fmt.Println(isUniqueString2(words))
	words = "dfjkalsu983"
	fmt.Println(isUniqueString1(words))
	fmt.Println(isUniqueString2(words))
}
