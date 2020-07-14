package stack

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	fmt.Println(IsValid("{}()[]"))
	fmt.Println(IsValid("({[]})"))
	fmt.Println(IsValid("((([[[{{{"))
	fmt.Println(IsValid("(())]]"))
}

func TestIsValid1(t *testing.T) {
	fmt.Println(IsValid1("{}()[]"))
	fmt.Println(IsValid1("({[]})"))
	fmt.Println(IsValid1("((([[[{{{"))
	fmt.Println(IsValid1("(())]]"))
}

func BenchmarkIsValid(b *testing.B) {
	for i := 1; i < b.N; i++ {
		IsValid("{{{{{[[[[[((((()))))]]]]]}}}}}")
	}
}

func BenchmarkIsValid1(b *testing.B) {
	for i := 1; i < b.N; i++ {
		IsValid1("{{{{{[[[[[((((()))))]]]]]}}}}}")
	}
}