package 面试题

import (
	"fmt"
	"testing"
)

func TestIsRegroup(t *testing.T) {
	s1 := "qqwweerrttyy1"
	s2 := "yyttwwrreeqq"

	fmt.Println(isRegroup(s1, s2))
}
