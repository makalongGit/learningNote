package 面试题

import (
	"fmt"
	"testing"
)

func TestReverString(t *testing.T){
	s := "djfaoisjglkasmdofgiwa"
	if s, ok := reverString(s);ok{

		fmt.Println(s)
	}
}
